package  interface_example

import "fmt"
import "math"

type UserAccess interface {
    getUsername() string
    getCurrentAccessKey() string
}

type Member struct {
    username string
    currentAccessKey string
    subscriptionType string
    expiredAt int
}

type Visitor struct {
    username string
    entryFee int
    currentAccessKey string
}

func Init() {
    fmt.Println("Init interface_example")

    var user1 UserAccess

    user1 = Member{
        username: "Surya",
        subscriptionType: "Gold",
        currentAccessKey: "I0F18J9JA",
        expiredAt: 1614410928,
    }
    fmt.Println("user1: ", user1.getUsername())

    // casting to Member struct
    fmt.Println("subscriptionType", user1.(Member).getSubscriptionType())

    var user2 UserAccess
    user2 = Visitor{
        username: "Respati",
        entryFee: 55000,
        currentAccessKey: "KAYG181NV",
    }

    fmt.Println("user2: ", user2.getUsername())
    fmt.Println("user2: ", user2.(Visitor).getEntryFee())

    var bangunRuang hitung = &kubus{4}

    fmt.Println("===== kubus")
    fmt.Println("luas      :", bangunRuang.luas())
    fmt.Println("keliling  :", bangunRuang.keliling())
    fmt.Println("volume    :", bangunRuang.volume())
}

func (m Member) getUsername() string {
    return m.username
}

func (m Member) getSubscriptionType() string {
    return m.subscriptionType
}

func (m Member) getCurrentAccessKey() string {
    return m.currentAccessKey
}

func (m Member) getExpiredAt() int {
    return m.expiredAt
}

func (v Visitor) getUsername() string {
    return v.username
}

func (v Visitor) getCurrentAccessKey() string {
    return v.currentAccessKey
}

func (v Visitor) getEntryFee() int {
    return v.entryFee
}

type hitung2d interface {
    luas() float64
    keliling() float64
}

type hitung3d interface {
    volume() float64
}

type hitung interface {
    hitung2d
    hitung3d
}

type kubus struct {
    sisi float64
}

func (k *kubus) volume() float64 {
    return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
    return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
    return k.sisi * 12
}



