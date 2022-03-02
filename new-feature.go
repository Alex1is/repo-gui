# Задание №1.

# 1 минута
one_minute = 60
# 1 час
one_hour = 3600
# 1 день
one_day = 86400
# 1 неделя
one_week = 604800

duration = int (input('Секунды: '))

#вывод в минутах:

if duration<one_minute:
    print ('{} сек'.format(duration))
#вывод в часах:

elif one_minute <= duration and one_hour > duration:
    my_minute=duration//one_minute
    my_second=duration%one_minute
    print ('{} мин {} сек'.format(my_minute,my_second));

#вывод в сутках:
elif duration >= one_hour and duration < one_day:
    my_hour=duration // one_hour
    duration=duration % one_hour
    my_minute = duration // one_minute
    my_second = duration % one_minute
    print ('{} час {} мин {} сек'.format(my_hour,my_minute,my_second));

#вывод в неделях:
elif duration >= one_day and duration < one_week:
    my_day = duration // one_day
    duration=duration % one_day
    my_hour = duration // one_hour
    duration = duration % one_hour
    my_minute = duration // one_minute
    my_second = duration % one_minute
    print('{} дн {} час {} мин {} сек'.format(my_day, my_hour, my_minute, my_second));

#Задание №2
# 1. Вычислить сумму тех чисел из этого списка, сумма цифр которых делится нацело на 7. Например, число «19 ^ 3 = 6859» будем включать в сумму, так как 6 + 8 + 5 + 9 = 28 – делится нацело на 7. Внимание: использовать только арифметические операции!
#создать список кубов нечётных чисел от 1 до 1000

my_list = []

for num in range(1,1001,2):
    my_list.append(num ** 3)

final_sum = 0
print(len(my_list))
#a
for num in my_list:
    check_sum = 0
    for check_num in str(num):
        check_sum += int(check_num)
    if check_sum % 7 == 0:
        final_sum += num
print(final_sum)
#b
final_sum = 0
for num in my_list:
    num += 17
    check_sum = 0
    for check_num in str(num):
        check_sum += int(check_num)
    if check_sum % 7 == 0:
        final_sum += num
print(final_sum)

#Задание №3.

#Реализовать склонение слова «процент» во фразе «N процентов». Вывести эту фразу на экран отдельной строкой для каждого из чисел в интервале от 1 до 100:

for i in range(100):
    new_str=str(i+1)
    new_list = list(new_str)
    if int(new_list[-1])==1 and i+1!=11:
        print('{} процент'.format(i + 1))
    elif int(new_list[-1])>1 and int(new_list[-1])<= 4:
        if  i+1> 10 and i+1<= 14:
            print('{} процентов'.format(i + 1))
        else:
            print('{} процента'.format(i + 1))
    else:
        print('{} процентов'.format(i + 1))
