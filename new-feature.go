#1. Написать генератор нечётных чисел от 1 до n (включительно), используя ключевое слово yield.

def odd_nums(max_value):
     n = 1
     while n <= max_value:
         yield n
         n += 2

odd_to_15 = odd_nums(15)

print(next(odd_to_15))
print(next(odd_to_15))
print(next(odd_to_15))
print(next(odd_to_15))
print(next(odd_to_15))

#3. Есть два списка:
#tutors = [
   # 'Иван', 'Анастасия', 'Петр', 'Сергей',
   # 'Дмитрий', 'Борис', 'Елена']
#klasses = ['9А', '7В', '9Б', '9В', '8Б', '10А', '10Б', '9А']
#Необходимо реализовать генератор, возвращающий кортежи вида (<tutor>, <klass>), например:
#('Иван', '9А')
#('Анастасия', '7В')

tutors = ['Иван', 'Анастасия', 'Петр', 'Сергей','Дмитрий', 'Борис', 'Елена']
klasses = ['9А', '7В', '9Б', '9В', '8Б', '10А', '10Б', '9А']

gen = ((tutor, klass) for tutor, klass in zip(tutors, klasses))
print(next(gen))

from itertools import zip_longest

gen = ((tutor, klass) for tutor, klass in zip_longest(tutors, klasses))
print(next(gen))

#4. Представлен список чисел. Необходимо вывести те его элементы, значения которых больше предыдущего.

src = [300, 2, 12, 44, 1, 1, 4, 10, 7, 1, 78, 123, 55]
new_list= [num for i, num in enumerate(src) if num > src[i - 1] and i !=0]
print(new_list)


#Подумайте, как можно сделать оптимизацию кода по памяти, по скорости.
#Представлен список чисел. Определить элементы списка, не имеющие повторений.
# Сформировать из этих элементов список с сохранением порядка их следования в исходном списке, например:
#src = [2, 2, 2, 7, 23, 1, 44, 44, 3, 2, 10, 7, 4, 11]
#result = [23, 1, 3, 10, 4, 11]

src = [2, 2, 2, 7, 23, 1, 44, 44, 3, 2, 10, 7, 4, 11]
print([x for x in src if src.count(x) == 1])

i = 0
new_list = []
while i < len(src):
    num = src[i]
    flag = True
    for el in src[i:]:
        if num == el:
            flag = False
    if flag:
        new_list.append(num)

    i += 1

