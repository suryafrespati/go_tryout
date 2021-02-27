package  interface_example

import "fmt"
// import "math"

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

