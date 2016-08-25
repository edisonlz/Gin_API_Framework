package user

import  (
    "github.com/astaxie/beego/orm"
    "fmt"
    "math/rand"
    "crypto/hmac"
    "time"
    "crypto/sha1"
    "encoding/hex"
)

type User struct {
    Id     int    `orm:"auto"`
    Name   string `orm:"size(100)"`
    Gender string `orm:"size(10)"`
    Phone  string `orm:"size(32)"`
    Email  string `orm:"size(32)"`
    Address string `orm:"size(64)"`
    Uuid string `orm:"size(128)"`
    Password string `orm:"size(128)"`
    Salt string `orm:"size(6)"`
}


func (user *User) CreateUser(
name string,
gender string,
phone string,
rawPassword string)  bool {

    //
    _, err := UserQueryByName(name)
    if err == nil{
        return false
    }
    o := orm.NewOrm()
    o.Using("default")
    //base user information
    user.Name = name
    user.Gender = gender
    user.Phone = phone

    //generator password
    salt := genSalt()
    password := genPassword(rawPassword, salt)
    user.Salt = salt
    user.Password = password

    o.Insert(user)
    return true
}

//generate salt
func genSalt() string{
    rnd := rand.New(
        rand.NewSource(time.Now().UnixNano()))
    vcode := fmt.Sprintf(
        "%06v", rnd.Int31n(1000000))
    return vcode
}

//gen password hmac key= salt , data rawPassowrd
func genPassword(rawPassowrd string, salt string) string {
    key := salt
    h := hmac.New(sha1.New, []byte(key))
    h.Write([] byte(rawPassowrd))
    password := hex.EncodeToString(
        h.Sum(nil))
    return password
}

//check user password
func (user *User) CheckPassword(rawPassword string) bool {
    password := genPassword(
        rawPassword, user.Salt)

    return password == user.Password
}

func UserQueryById(uid int) (user User) {

    o := orm.NewOrm()
    u := User{Id: uid}

    err := o.Read(&u)

    if err == orm.ErrNoRows {
        fmt.Println("查询不到")
    } else if err == orm.ErrMissPK {
        fmt.Println("找不到主键")
    } else {
        fmt.Println(u.Id, u.Name)
    }

    return u
}

func UserQueryByName(name string) (User, error) {
    var user User

    o := orm.NewOrm()
    qs := o.QueryTable("user")

    err := qs.Filter("Name", name).One(&user)
    fmt.Println(err)
    if err == nil {
        fmt.Println(user.Name)
        return user,nil
    }
    return user, err
}

func UserList() (users []User) {

    o := orm.NewOrm()
    qs := o.QueryTable("user")

    var us []User
    cnt, err :=  qs.Filter("id__gt", 0).OrderBy("-id").Limit(10, 0).All(&us)
    if err == nil {
        fmt.Printf("count", cnt)
        // for _, u := range us {
        //     fmt.Println(u)
        // }
    }
    return us
}

type Post struct {
    Id    int    `orm:"auto"`
    Title string `orm:"size(100)"`
    User  *User  `orm:"rel(fk)"`
}
