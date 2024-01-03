import "fmt"
func canVisitAllRooms(rooms [][]int) bool {
    keys := []int{}
    visited := map[int]bool{}
    
    visited[0] = true
    for _, key := range rooms[0] {
        keys = append(keys, key)
        visited[key] = true
    }

    for len(keys) > 0 {
        newKeys := []int{}
        for _, room := range keys {
            for _, key := range rooms[room] {
                if _, ok := visited[key]; !ok {
                    newKeys = append(newKeys, key)
                    visited[key] = true
                }                
            }
        }

        keys = newKeys
    }

    return len(visited) == len(rooms)
}
