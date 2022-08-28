# `tsc`
> 一个将 `json` 数据结构, 转换为 `typescript` 数据模型的转换器的项目
>


[English](./README.md) | 中文



## 1.用法

- `tsc -d '${json.string.content}'`

    - ```shell
  $ tsc -d '"{\"name|required\":\"photowey\",\"age|readonly\":18}"'

  # json 内容需要转义
    ```

- `tsc --data '${json.string.content}'`

    - ```shell
  $ tsc --data '"{\"name|required\":\"photowey\",\"age|readonly\":18}"'
    ```

- `tsc in ${json.file}`

    - > 现在还不支持: `Linux` 输入重定向，通过 **in** 子命令替代 **<**
      >
      > `tsc < ${json.file} `

    - ```shell
  $ tsc in example.json
    ```
  
  - ```shell
  $ tsc in example1.json example2.json
    ```

    - ```shell
  $ tsc in example1.json ... exampleN.json
    ```



## 2.输出示例:

### 2.1.`json` 文件

```json
{
  "name|required": "photowey",
  "age|readonly": 18,
  "balance": 10.24,
  "boy": true,
  "address": "chongqing",
  "hobby": [
    {
      "name": "badminton",
      "description": "badminton"
    }
  ],
  "university": {
    "name": "cqjtu",
    "address": "ertang"
  }
}
```

```shell
# 注意
# 我们可以通过 `required` 和 `readonly` 来修饰 `json` 数据结构的 `key`。
# 从而, 达到对 `TS` 数据模型属性的控制
```

### 2.2.`typescript` 输出模型

> 默认的数据模型的名称为: `DataModel`
>
> > 由于 `GO` 在处理 `JSON` 内容的时候会将其转化为 `map[string]any` 也就是一个无序数据结构
> >
> > > 这将导致: 每次输出的数据模型的属性的顺序不太一样
> >
> > 现在支持的数据类型:
> >
> > - `string`
> > - `number`
> > - `boolean`

```typescript
// ---------------------------------------------------------------- example.json
export interface DataModel {
    readonly age?: number // 18
    balance?: number // 10.24
    boy?: boolean // true
    address?: string // chongqing
    Hobby?: Hobby[] // Hobby
    University?: University // University
    name: string // photowey
}

export interface Hobby {
    name?: string // badminton
    description?: string // badminton
}

export interface University {
    address?: string // ertang
    name?: string // cqjtu
}
```

