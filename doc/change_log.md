# Change Log

**v0.9.0**
- list remove コマンドを追加
- list member コマンドを追加

**v0.8.1**
- go version 1.11 -> 1.13
- モジュールの管理を dep から go mod に変更
- モジュールのバージョンアップ
- show list, list addコマンドで使用していた anaconda を go-twitter に差し替え

**v0.8.0**
- list add コマンドを追加

**v0.7.0**
- show list コマンドを追加

**v0.6.0**
- show profile コマンドを追加
    - --by-id オプション: screen nameではなく、user idで検索
- delete, unlike (unf) コマンドを追加

**v0.5.0**
- show home コマンドに下記オプションを追加
    - screen name 引数: ログインユーザー以外のUserTimelineを取得可能にした
    - --page オプション: limit数のツイートを page回分取得
    - --since-id オプション: since-idより新しいツイートまで取得
    - --max-id オプション: max-idと同じか古いツイートから取得
    - --retry オプション: 通信エラー時、再チャレンジする回数。初期値 3 回
- show friend, show follower コマンドを追加
    - --limit, --page オプション: limit数のツイートを page回分取得
    - --all オプション: friend, followerをすべて取得
    - --as-id オプション: user idの一覧を取得
    - --retry オプション: 通信エラー時、再チャレンジする回数。初期値 3 回
- show home, show tweet, show favorite コマンドの help 中 limit オプションに最大値を記載

**v0.4.0**
- show tweet コマンドに下記オプションを追加
    - screen name 引数: ログインユーザー以外のUserTimelineを取得可能にした
    - --page オプション: limit数のツイートを page回分取得
    - --since-id オプション: since-idより新しいツイートまで取得
    - --max-id オプション: max-idと同じか古いツイートから取得
    - --retry オプション: 通信エラー時、再チャレンジする回数。初期値 3 回
- show favorite コマンドを追加

**v0.3.1**
- show home コマンドで show tweetコマンドのフラグを再定義する不具合を修正
- tweet (say) コマンドの --id オプションの名称が誤っていた不具合を修正

**v0.3.0**
- tweet, say コマンドに --id オプションを追加。リプライ機能
- like (f), retweet (rt), quote (qt)コマンドを追加
- ツイート表示を調整

**v0.2.0**
- コマンドをcobraベースに変更
- gatapi tweet, say, show home, show tweetコマンドを追加

**v0.1.0**
- gatapi コマンドの作成
- ログインユーザーのUser Timelineを10件取得するだけ
