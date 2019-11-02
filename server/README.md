# Torch Server

## @nanagami 謹製CSVもどきをgormのクエリにしよう

正規表現するやつ。

Search :

```regexp
(.+?)[ \t]*([\d\.]+),([\d\.]+)\n\t+([\d\.]+),([\d\.]+)
```

Replace : 

```regexp
db.Create(&types.AreaInfo{Name: "$1", LeftUpX: $2, LeftUpY: $3, RightBottomX: $4, RightBottomY: $5})
```