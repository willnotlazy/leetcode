package questions

func OpenLock(deadends []string, target string) int {
    queue := []string{"0000"}
    dead_map := make(map[string]struct{})
    visited := make(map[string]struct{})

    for _, deadend := range deadends {
        dead_map[deadend] = struct{}{}
    }
    var step int
    visited["0000"] = struct{}{}
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i ++ {
            cur := queue[i]
            if _, ok := dead_map[cur]; ok {
                continue
            }

            if cur == target {
                return step
            }
            
            for j := 0; j < 4; j ++ {
                cur_up := up(cur, j)
                cur_down := down(cur, j)
                if _, ok := visited[cur_up]; !ok {
                    visited[cur_up] = struct{}{}
                    queue = append(queue, cur_up)
                }

                if _, ok := visited[cur_down]; !ok {
                    visited[cur_down] = struct{}{}

                    queue = append(queue, cur_down)
                }
            }
        }

        step ++
        queue = queue[size:]
    }

    return -1
}

func up(s string, pos int) string {
    b := []byte(s)
    if b[pos] == '9' {
        b[pos] = '0'
    } else {
        b[pos] ++
    }

    return string(b)
}

func down(s string, pos int) string {
    b := []byte(s)
    if b[pos] == '0' {
        b[pos] = '9'
    } else {
        b[pos] --
    }

    return string(b)
}