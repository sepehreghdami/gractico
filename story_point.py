num = int(input())

fib_list=[1,2]

while fib_list[-1] <= num:

    fib_list.append(fib_list[-1]+fib_list[-2])


if num in fib_list:
    print(1)


else:
    sum = fib_list[-2]
    count = 1
    for i in fib_list[:-2][::-1]:
        if sum+i <= num:
            sum += i
            count += 1
            if sum == num:
                break

    print(count)    