func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0);

    for i:=0; i<len(asteroids); i++ {
        ast := asteroids[i];

        if len(stack) == 0 {
            stack = append(stack, ast);
            continue;
        }

        if stack[len(stack)-1] > 0 && ast > 0 {
            stack = append(stack, ast);
            continue;
        }

        if stack[len(stack)-1] < 0 && ast < 0 {
            stack = append(stack, ast);
            continue;
        }

        if stack[len(stack)-1] < 0 && ast > 0 {
            stack = append(stack, ast);
            continue;
        }

        if stack[len(stack)-1] > 0 && ast < 0 {
            for stack[len(stack)-1] <= -ast {
                if stack[len(stack)-1] < 0 {
                    stack = append(stack, ast);
                    break;
                }
                
                if stack[len(stack)-1] == -ast {
                    stack = stack[:len(stack)-1];
                    break;
                }

                stack = stack[:len(stack)-1];

                if len(stack) == 0 {
                    stack = append(stack, ast);
                    break;
                }
            }
            continue;
        }
    }

    return stack;
}
