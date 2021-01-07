package models

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id int64 `orm:"auto" json:"id,omitempty"`
	UserBody
	UserInfoBody
	Profile *Profile `orm:"rel(one)" json:"profile"` // OneToOne relation
	TimeStamp
}

type UserBody struct {
	Email    string `orm:"size(128)" json:"email"`
	Password string `orm:"size(128)" json:"password"`
}

type UserResponse struct {
	Id int64 `orm:"auto" json:"id,omitempty"`
	UserLoginResponseBody
	UserInfoBody
	Profile *Profile `orm:"rel(one)" json:"profile"` // OneToOne relation
	TimeStamp
}

type UserLoginResponseBody struct {
	Email    string `orm:"size(128)" json:"email"`
	Password string `orm:"size(128)" json:"-"`
}

type UserInfoBody struct {
	Name string `orm:"size(128)" json:"name" example:"山田太郎"`
}
type Profile struct {
	Id int64 `orm:"auto" json:"id,omitempty"`
	ProfileBody
	TimeStamp
}

type ProfileBody struct {
	Age int64 `orm:"size(128)" json:"age"`
}

type Token struct {
	Id    int64  `orm:"auto" json:"id,omitempty"`
	Token string `json:"token"`
	User  *User  `orm:"rel(one)" json:"user"` // OneToOne relation
	TimeStamp
	ExpiredAt time.Time `json:"-"`
}

type AccessToken struct {
	Id    int64  `orm:"auto" json:"id,omitempty"`
	Token string `json:"token"`
	User  *User  `orm:"rel(one)" json:"user"` // OneToOne relation
	TimeStamp
	ExpiredAt time.Time `json:"-"`
}

var (
	UserAlreadyExistsError = errors.New("user_already_exists")
)

func (user *User) CheckPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	hash := hashAndSalt(m.Password)
	m.Password = hash

	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	count, err := qs.Filter("Email", m.Email).Count()
	if err != nil {
		return
	}
	if count != 0 {
		return 0, UserAlreadyExistsError
	}
	profileId, err := o.Insert(m.Profile)
	m.Profile.Id = profileId
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CreateToken(m *User) string {
	o := orm.NewOrm()
	token := hashAndSalt(string(m.Id))
	expired_at := time.Now().AddDate(0, 0, 30)

	t := new(Token)
	if err := o.QueryTable("token").Filter("User", m).One(t); err != orm.ErrNoRows {
		t.Token = token
		t.ExpiredAt = expired_at
		o.Update(t)
	} else {
		t.Token = token
		t.User = m
		t.ExpiredAt = expired_at
		o.Insert(t)
	}
	return token
}

func CreateAccessToken(m *User) string {
	o := orm.NewOrm()
	token := hashAndSalt(string(m.Id))
	expired_at := time.Now().Add(3 * time.Hour)

	t := new(AccessToken)
	if err := o.QueryTable("access_token").Filter("User", m).One(t); err != orm.ErrNoRows {
		t.Token = token
		t.ExpiredAt = expired_at
		o.Update(t)
	} else {
		t.Token = token
		t.User = m
		t.ExpiredAt = expired_at
		o.Insert(t)
	}
	return token
}

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
