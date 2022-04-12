#1. Написать скрипт, создающий стартер (заготовку) для проекта со следующей структурой папок:
#|--my_project
#   |--settings
#   |--mainapp
#   |--adminapp
#   |--authapp
#Примечание: подумайте о ситуации, когда некоторые папки уже есть на диске (как быть?);
#как лучше хранить конфигурацию этого стартера, чтобы в будущем можно было менять имена папок под конкретный проект;
#можно ли будет при этом расширять конфигурацию и хранить данные о вложенных папках и файлах (добавлять детали)?
 import os

 pattern = {'my_project': ['settings', 'mainapp', 'adminapp', 'authapp']}
 for root, folders in pattern.items():
     if os.path.exists(root):
         print(root, 'существует')
     else:
         for folder in folders:
             cur_dir = os.path.join(root, folder)
             os.makedirs(cur_dir)


#3. Создать структуру файлов и папок, как написано в задании 2 (при помощи скрипта или «руками» в проводнике). Написать скрипт, который собирает все шаблоны в одну папку templates, например:
#|--my_project
#   ...
#  |--templates
#   |   |--mainapp
#   |   |  |--base.html
#   |   |  |--index.html
#   |   |--authapp
#   |      |--base.html
#   |      |--index.html
#Примечание: исходные файлы необходимо оставить; обратите внимание, что html-файлы расположены в родительских папках
#(они играют роль пространств имён); предусмотреть возможные исключительные ситуации;
#это реальная задача, которая решена, например, во фреймворке django.


 import os
 import shutil

 folder = r'my_project2'
 files = []
 my_dir = os.path.join(folder, 'templates')
 if not os.path.exists(my_dir):
     os.mkdir(my_dir)


 for r, d, f in os.walk(folder):
     for file in f:
         if '.html' in file:
             files.append(os.path.join(r, file))
 for path in files:
     folder = os.path.join(my_dir, os.path.basename(os.path.dirname(path)))
     if not os.path.exists(folder):
         os.mkdir(folder)
     save_path = os.path.join(folder, os.path.basename(path))
     shutil.copy(path, save_path)

#4. Написать скрипт, который выводит статистику для заданной папки в виде словаря, в котором ключи — верхняя граница размера файла (пусть будет кратна 10), а значения — общее количество файлов (в том числе и в подпапках), размер которых не превышает этой границы, но больше предыдущей (начинаем с 0), например:
#    {
#      100: 15,
#      1000: 3,
#      10000: 7,
#      100000: 2
#    }
#Тут 15 файлов размером не более 100 байт; 3 файла больше 100 и не больше 1000 байт...
#Подсказка: размер файла можно получить из атрибута .st_size объекта os.stat.



 import os
 files = []
 for r, d, f in os.walk('./'):
     for file in f:
         file_path = os.path.join(r, file)
         files.append(os.stat(file_path).st_size)
 max_size = max(files)
 i = 1
 out_dict = {}

 for _ in range(len(str(max_size))):
     i *= 10
     out_dict[i] = 0
 for file in files:
         out_dict[10 ** len(str(file))] += 1


 print(out_dict)