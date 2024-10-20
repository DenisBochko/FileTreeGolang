# FileTreeGolang
FileTree - программа для вывода дерева файлов по указанному пути. </br>
Чтобы воспользоваться, необходимо выполнить в дериктории с файлом команду: </br>
```go run main.go . -f```
- флаг ```-f``` - включает вывод файлов, без него будут выводиться только директории </br>
- ```.``` - указывает на путь </br>
## Пример вывода: </br>
```
.
├───main.go
├───picturesForTest
│   ├───image.png
│   └───image1.png
├───README.md
└───testdata
    ├───project
    │   ├───file.txt
    │   └───gopher.png
    ├───static
    │   ├───a_lorem
    │   │   ├───dolor.txt
    │   │   ├───gopher.png
    │   │   └───ipsum
    │   │       └───gopher.png
    │   ├───css
    │   │   └───body.css
    │   ├───empty.txt
    │   ├───html
    │   │   └───index.html
    │   ├───js
    │   │   └───site.js
    │   └───z_lorem
    │       ├───dolor.txt
    │       ├───gopher.png
    │       └───ipsum
    │           └───gopher.png
    ├───zline
    │   ├───empty.txt
    │   └───lorem
    │       ├───dolor.txt
    │       ├───gopher.png
    │       └───ipsum
    │           └───gopher.png
    └───zzfile.txt
```
