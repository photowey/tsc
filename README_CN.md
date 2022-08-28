# `tsc`
> 一个将 `json` 数据结构, 转换为 `typescript` 数据模型的转换器的项目
>


[English](./README.md) | 中文

## 1.用法

### 1.1.`tsc --data`

1.1.1.`tsc -d '${json.string.content}'`

```shell
$ tsc -d '"{\"name|required\":\"photowey\",\"age|readonly\":18}"'
    
# json 需要转义为 json 字符串
```

1.1.2.`tsc --data '${json.string.content}'`

```shell
$ tsc --data '"{\"name|required\":\"photowey\",\"age|readonly\":18}"'
```

### 1.2.`tsc in`

> 现在还不支持: `Linux` 输入重定向, 故采用 **`in`** 子命令来替换 **`<`**

1.2.1.`tsc in ${json.file}`

```shell
$ tsc in example.json

$ tsc in example1.json example2.json

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

