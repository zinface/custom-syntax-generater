# custom-syntax-generater

A simple syntax custom ideator

## Tr Series Components

### 1. manpage-tr  - 语义化转换生成标准 man 语法手册

> 将初步基本形式的 man.tr 转换成 man 手册语法

1. 编辑：manpage-tr.tr
    ```
    $$comment syntax:   P  N             T                   V        T
    $$comment e.g: .TH MAN 1 2021-03-25 15:43:59 +0800 CST 2.8.5 手册分页显示工具
    
    $$title mampage-tr 1 "$$date" 0.0.1 "ManPage 生成工具"
    $$subtitle NAME
        manpage-tr - manpage 生成、定制转换工具

    $$subtitle tr 组件系列
        manpage-tr 是该系列的始祖，通过该系列有以下各种工具

        $$subcontent manpage-tr 
            $$cbegin
                manpage 生成、定制转换工具
            $$cend
        $$subcontent markdown-tr 
            $$cbegin
                markdown 生成、定制转换工具
            $$cend
        $$subcontent html-tr
            $$cbegin
                html 生成、定制转换工具
            $$cend
        $$subcontent javadoc-tr 
            $$cbegin
                javadoc
            javadoc 生成、定制转换工具
   ```


2. 执行 `manpage-tr manpage-tr.tr` 后生成 `manpage-tr.1`

    ```tr
    mampage-tr(1)                   ManPage 生成工具                  mampage-tr(1)

    NAME
        manpage-tr - manpage 生成、定制转换工具

    tr 组件系列
        manpage-tr 是该系列的始祖，通过该系列有以下各种工具

        manpage-tr
                manpage 生成、定制转换工具

        markdown-tr
                markdown 生成、定制转换工具

        html-tr
                html 生成、定制转换工具

        javadoc-tr
                javadoc 生成、定制转换工具

    0.0.1                    2021-03-25 15:43:59 +0800 CST            mampage-tr(1)
    ```
