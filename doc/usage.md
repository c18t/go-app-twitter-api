# Usage

## お気に入りユーザーランキング
```sh
$ gatapi show favorite <screen name> -l 200 -p 100 | grep -p '@\w\+:' | awk -F '[@:]' '{print $2}' | sort | uniq -c | sort -nr
```

## お気に入りユーザーランキングから2like以上したユーザーをlist化
```sh
$ gatapi list member --list-id <list id> -a | grep -p '@\w\+:' | awk -F '[@:]' '{print $2}' | sort > member-b.txt
$ gatapi show favorite <screen name> -l 200 -p 100 | grep -p '@\w\+:' | awk -F '[@:]' '{print $2}' | sort | uniq -c | sort -nr | awk -F ' ' '{if ($1 > 1) print $NF}' | sort > member-a.txt
$ comm -23 member-b.txt member-a.txt | xargs -n 100 gatapi list remove --list-id <list id>
$ comm -13 member-b.txt member-a.txt | xargs -n 100 gatapi list add --list-id <list id>
$ rm member-b.txt member-a.txt
```