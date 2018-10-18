# Usage

## お気に入りユーザーランキング
```sh
$ gatapi show favorite <screen name> -l 200 -p 100 | grep -p '@\w\+:' | awk -F '[@:]' '{print $2}' | sort | uniq -c | sort -nr
```

## お気に入りユーザーランキングから2like以上したユーザーをlist化
```
$ gatapi show favorite <screen name> -l 200 -p 100 | grep -p '@\w\+:' | awk -F '[@:]' '{print $2}' | sort | uniq -c | sort -nr | awk -F ' ' '{if ($1 > 1) print $NF}' | xargs gatapi list add 
```