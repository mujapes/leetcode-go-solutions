func recoverOrder(order []int, friends []int) []int {
    friendsOrder := make([]int, len(friends))
    nextFriendIndex := 0
    for _, o := range order {
        for _, f := range friends {
            if o == f {
                friendsOrder[nextFriendIndex] = f
                nextFriendIndex++
            }
        }
    }
    return friendsOrder
}

//Runtime: 0ms
//Memory: 5.27MB