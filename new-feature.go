#1. Написать функцию email_parse(<email_address>),
#которая при помощи повторяющегося обращения из найденного имени пользователя и поиска домена из адреса электронной почты
#и возвращает их в виде словаря. Если адрес недействителен, сбросьте ошибку ValueError.
#Пример:
#>>> email_parse('someone@geekbrains.ru')
#{'username': 'someone', 'domain': 'geekbrains.ru'}
#>>> email_parse('someone@geekbrainsru')
#Traceback (most recent call last):
#  File "<stdin>", line 1, in <module>
#  ...
#    raise ValueError(msg)
#ValueError: wrong email: someone@geekbrainsru
#Примечание: подумайте о возможных ошибках в адресе и постарайтесь зафиксировать их в проявлениях;
#имеет ли смысл использовать функцию re.compile()?

 import re
 EMAIL = re.compile(r'([a-z0-9]+)@([a-z0-9]+\.[a-z]+)')


 def email_parse(email):
     found_info = EMAIL.findall(email)
     if found_info and found_info[0]:
         name, addr = found_info[0]
     else:
         raise ValueError(f'wrong email: {email}')
     return name, addr

 print(email_parse('someone@geekbrains.ru'))
 print(email_parse('someone@geekbrainsru'))

#3. Написать декоратор для логирования типовых позиционных аргументов функций, например:
#def type_logger...
#    ...


#@type_logger
#def calc_cube(x):
 #  return x ** 3

#>>> a = calc_cube(5)
#5: <class 'int'>
#Примечание: если аргументов несколько - выводить данные о каждой через запятую;
# Можете ли вы указать тип значения функции? Можно ли решить проблему именованных аргументов?
# Сможете ли вы замаскировать работу декоратора? Можно ли вывести имя функции, например, в виде:
#>>> a = calc_cube(5)
#calc_cube(5: <class 'int'>)

 from functools import wraps


 def type_logger(func):
     @wraps(func)
     def wrapper(*args):
         """qwe"""
         for arg in args:
             print(f'{func.__name__}({arg}: {type(arg)})', end=', ')
         return func(*args)

     return wrapper


 @type_logger
 def calc_cube(*args):
     """this shows only with 'from functools import wraps'"""
     return list(map(lambda x: x ** 3, args))

 a = calc_cube(5, 3)
 print(a)
 print(calc_cube.__name__)
 print(calc_cube.__doc__)

#4. Написать декоратор с аргументом-функцией (обратный вызов), 
# считать допустимыми входные значения функций и выбрасывать исключение ValueError, если что-то не так, например:
#def val_checker...
#    ...


#@val_checker(lambda x: x > 0)
#def calc_cube(x):
#   return x ** 3


#>>> a = calc_cube(5)
#125
#>>> a = calc_cube(-5)
#Traceback (most recent call last):
#  ...
 #   raise ValueError(msg)
#ValueError: wrong val -5


from functools import wraps


def val_checker(decorator_check_func):
    def _val_checker(func_calc_cube):
        @wraps(func_calc_cube)  
        def wrapped(x):
            if decorator_check_func(x):
                return func_calc_cube(x)
            else:
                raise ValueError(x)

        return wrapped
    return _val_checker


@val_checker(lambda x: x > 0)
def calc_cube(x):
    """calc_cube desc"""
    return x ** 3


a = calc_cube(-5)
print(a)
print(calc_cube.__name__)
print(calc_cube.__doc__)

