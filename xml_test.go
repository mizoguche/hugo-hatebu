package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	as := assert.New(t)
	rss, err := Parse(testXml)
	as.Nil(err)
	as.Equal(15, len(rss.Channel.Items))
	as.Equal("Pull RequstのレビューにLGTM画像はっつけるChrome拡張つくった", rss.Channel.Items[0].Title)
	as.Equal(3, len(rss.Channel.Items[0].Categories))
	as.Equal("JavaScript", rss.Channel.Items[0].Categories[0])
}

func TestItem_Date(t *testing.T) {
	as := assert.New(t)
	rss, err := Parse(testXml)
	as.Nil(err)
	d := rss.Channel.Items[0].Date()
	as.Equal(time.Date(2017, 4, 10, 0, 0, 0, 0, time.FixedZone("", 3600)), d)
}

func TestRSS_Latest(t *testing.T) {
	as := assert.New(t)
	rss, err := Parse(testXml)
	as.Nil(err)
	as.Equal("Pull RequstのレビューにLGTM画像はっつけるChrome拡張つくった", rss.Latest().Title)
}

const testXml = `
<?xml version="1.0" encoding="utf-8" standalone="yes" ?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
    <channel>
        <title>mizoguche.info on mizoguche.info</title>
        <link>http://localhost:1313/</link>
        <description>Recent content in mizoguche.info on mizoguche.info</description>
        <generator>Hugo -- gohugo.io</generator>
        <lastBuildDate>Mon, 10 Apr 2017 00:00:00 +0000</lastBuildDate>
        <atom:link href="/" rel="self" type="application/rss+xml" />

        <item>
            <title>Pull RequstのレビューにLGTM画像はっつけるChrome拡張つくった</title>
            <link>http://localhost:1313/2017/04/approve_with_lgtm_image/</link>
            <pubDate>Mon, 10 Apr 2017 00:00:00 +0100</pubDate>

            <guid>http://localhost:1313/2017/04/approve_with_lgtm_image/</guid>
            <description> こんなん作った。
 mizoguche/approve-pr-with-image: Chrome Extension to approve pull request with LGTM image. Approve Pull Request with image - Chrome ウェブストア  動機  LGTM画像の貼り付けにlgtm.inからランダムに画像取得して貼り付けるChrome拡張使ってたけどもっさりした動作でストレスフルだった 以下のような思いつき  Local StorageにURL保存しといたらすぐ貼り付けられるのでは いいLGTM画像見つけたら右クリックでURL保存できるようにすると捗りそう  Reactを触っときたかった  感想  LGTM画像の貼り付けのストレスがだいぶ軽減された ついでにApproveのラジオボタンに自動的にチェック入るようにしたので捗る React/Redux/redux-observableは仕様に対してかなりオーバーテクノロジーでjQueryで十分という感じだっだが業務に活かされたので良しとする  React/Redux/redux-observableに関してはあとでまとめたい   </description>

            <category>JavaScript</category>

            <category>Chrome拡張</category>

            <category>Approve Pull Request With LGTM Image</category>

        </item>

        <item>
            <title>Goglandでgoimportsするショートカットをつくる</title>
            <link>http://localhost:1313/2017/01/intellij_goimports/</link>
            <pubDate>Fri, 06 Jan 2017 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2017/01/intellij_goimports/</guid>
            <description> Gogland使い始めてIntelliJ IDEAに設定してたgoimportsのショートカットどうやったか忘れてた。
goimportsを使ってIntelliJ IDEA でGoのコードのimportを手軽に整理する - Qiitaを参考にしながらいろいろやった。
TL; DR  External Toolsからコマンドを実行する 追加したExternal Toolsにショートカットキーを設定できる  1. スクリプトを作成して適当な場所に置いて実行権限を与える  goimportsのついでにgofmt, golint, go vetもするようにした。
2. External Toolsの設定 2.1. Preferences &amp;gt; Tools &amp;gt; External Toolsから追加 2.2. スクリプトをToolとして追加する 3. ショートカットキーを設定する 参考  goimportsを使ってIntelliJ IDEA でGoのコードのimportを手軽に整理する - Qiita  </description>

            <category>Gogland</category>

            <category>IntelliJ IDEA</category>

        </item>

        <item>
            <title>macOS SierraにアップデートしてHammerspoonでCommandキーにかなと英数を割り当てた</title>
            <link>http://localhost:1313/2017/01/hammerspoon_for_sierra/</link>
            <pubDate>Wed, 04 Jan 2017 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2017/01/hammerspoon_for_sierra/</guid>
            <description> macOS SierraにアップデートしたらKarabinerが使えなくなるという話だったのでYosemiteからのアップデートを先延ばしにしてたところ、XcodeをアップデートのためにOSをアップデートする必要が出たのでSierraにアップデートし、KarabinerからHammerspoonに乗り換えた。
TL; DR  初めてのLuaをふんわり書いたら動いた 今のところKarabinerでやってることは実現できてるのでSierraでも問題を感じない  Hammerspoon Karabiner 使えない対策: Hammerspoon で macOS の修飾キーつきホットキーのキーリマップを実現する - QiitaとかRebuild: 166: Self-Confident Neural Network (naoya)あたりでHammerspoonの存在を把握してたのでとりあえず乗り換えた。
ダウンロード・インストール・起動してOpen ConfigをクリックするとLuaスクリプトの編集ができ、Reload Configをクリックすると変更が反映される。
やりたいこと  ESCで英数入力にする 左Commandで英数入力にする 右Commandでかな入力にする  init.lua  参考  Switch to specific layout on command left key press only? · Issue #1039 · Hammerspoon/hammerspoon Hammerspoon docs: hs.eventtap module 英数キーとかかなキーとかを同時押しの時だけAltキーにするHammerspoonの設定  まとめ Lua  初めて書いたけど、リファレンスとか全く見ずに見よう見まねでやりたいことの実現ができた  Hammerspoon  ちゃんとAPIリファレンスがあって、カスタマイズには困らなさそう  カスタマイズは必要最低限に抑えたいとは思っている   </description>

            <category>Hammerspoon</category>

            <category>Sierra</category>

            <category>Lua</category>

        </item>

        <item>
            <title>SHIROBAKOで学ぶプロジェクトマネジメント #2 長時間労働のデメリット</title>
            <link>http://localhost:1313/2016/12/shirobako_2/</link>
            <pubDate>Thu, 22 Dec 2016 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2016/12/shirobako_2/</guid>
            <description> SHIROBAKO #02 あるぴんはいます！  ４話の制作状況が逼迫し、慌てるあおい。３話作画監督の遠藤にヘルプを直訴し、やり取りの末、なんとか承諾を得る。だが、一難去ってまた一難−−。4話のダビングに出向いたあおいは、監督と演出家のリテイクに対する応酬に胸を騒がせることに。声優の演技に触発された監督の注文は、音付けから次第にエスカレートし、絵に対するものへと矛先を変えていく。そしてあおいは、監督から発せられた、信じられない一言を耳にした。
 主な出来事  瀬川さんが過負荷により発熱して休養 作画リテイク  TL; DR  長時間労働を肯定すると問題が見えなくなる 手戻りが発生したら、プロセスの改善によって次回以降で同じようなことを避けるべき  長時間労働のデメリット パフォーマンスの低下  明らかになったのは、１日に動ける時間の長さや、体の部位を動かせる回数は人それぞれに決まっていて、「昨日の遅れを今日取り戻そう」と長時間働いても、こなせる仕事量は実はほとんど増えないということです。
発注にノーと言わぬ上司、長時間労働の要因　産業医指摘：朝日新聞デジタル
 劇中でも瀬川さんが過労による発熱によりダウンし、トータルで見たときの成果として長時間労働した分成果が増えたかは疑問である。もちろん締め切りに間に合わせるという意味での成果はあったが、ダウンした分は後ろのスケジュールにしわ寄せが行ってるのでトータルでみるとどうなんだろうか、という。
問題が見えなくなる コストとリターンのトレードオフを考慮した上で、クオリティを上げるための作画リテイクは選択肢として常に持つべき(劇中で手戻りによるコスト増の話が出ないところに闇を感じる部分もあるが)。
しかし、「徹夜して頑張ったらいいものができた！」と暗に長時間労働を肯定するような描き方をされてる印象を受けた。
 「時間内に終わらなければ残業すればいい」という考え方で対処していると、なぜ仕事が終わらないかという理由がわからず、…常に同じ問題が繰り返し発生し続けることになります。 つまり、問題を顕在化し改善する絶好の機会が、残業によって奪われてしまうのです。
「残業ゼロの仕事力」吉越浩一郎(日本能率協会マネジメントセンター、2007年) p.17
 リテイクによってスケジュールが圧迫されたなら、事後に、
 リテイクの原因をつきとめる 根本的な解決策を考える 再発を防止する  ことで、プロセスを改善して同じクオリティの成果をより少ない時間で出せるようにする必要がある。
木こりのジレンマにハマると組織として競争力を失っていく。プロセスの改善ができない組織はいずれ淘汰されるのではないかという危機感から、イテレーションの振り返りを行っているのであります。
まとめ  長時間労働は生産性を下げる 長時間労働は問題点を覆い隠す プロセスの改善によって成果物のクオリティを上げないと死ぬかもしれないの怖い   SHIROBAKO Blu-ray プレミアムBOX vol.1(初回仕様版) 著者: ワーナー・ブラザース・ホームエンターテイメント 発売日: 2016-11-23    新装版 「残業ゼロ」の仕事力 著者: 吉越 浩一郎 発売日: 2011-01-28   </description>

            <category>プロジェクトマネジメント</category>

            <category>SHIROBAKOで学ぶプロジェクトマネジメント</category>

        </item>

        <item>
            <title>SHIROBAKOで学ぶプロジェクトマネジメント #1 怒る必要ない</title>
            <link>http://localhost:1313/2016/12/shirobako_1/</link>
            <pubDate>Thu, 01 Dec 2016 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2016/12/shirobako_1/</guid>
            <description> SHIROBAKOがプロジェクトマネジメントの教科書なので見ろということで、Blu-ray BOXとPS3がヤマトで送られてきたので全部見た。
1話ずつまとめてく。
SHIROBAKO #01 明日に向かって、えくそだすっ！
 「どんどんドーナツ！　どんと行こう！」。上山高校アニメーション同好会のあおい、絵麻、しずか、美沙、みどりは、いつか共に商業アニメーションを作ろうとドーナツに誓った。そして、二年半後。かつて夢を掲げた手には無骨なハンドルが握られている。あおいは、アニメーションの制作進行として今日も精一杯駆け回っていた。回収からスタジオに戻り、見回すと人の気配が消えている。不審に思った彼女が会議室のドアを開くと−−。
 プロジェクトマネジメント上のヤバいこと 高梨がスケジュールの遅延を報告していなかった 割とよく聞く話。12月完了予定のプロジェクトが数ヶ月遅れることを11月に発表するのも、下から上がってくるスケジュールの予実が粉飾されてたから起こることだと推測する。ちゃんと把握してたらもっと前に延期しなければいけないことがわかるし、ダメージコントロールの観点からも早めに発表したほうがいいので直前に発表するメリットがない。
たぶん内部スケジュール(実際のスケジュール)と外部スケジュール(外部に報告する見栄えのいいスケジュール)が存在してる。
原因と対策 正直にスケジュールを報告 しない インセンティブを与えない スケジュールの遅れを報告したときに怒ると、「怒られるから報告しない」というインセンティブを生む(最終話までに「矢野さんは怒るからイヤ」という趣旨の発言が高梨から出てた)。つまり、怒る・詰めるなどをするマネージャーは自分で自分の首を締めてる。
怒る・詰めるなどをしても何もメリットはないし問題は解決しない。怒られて問題の再発が防止できるなら、マネージャーという職業はいらない。
問題を解決して対策を考えて再発を防止するをということに責任を負ってるハズなのに、その責任を部下に丸投げするパターンは割とよく聞く。
敵対しない 怒った時点で、作業者とマネージャーが敵対関係になるのがマズい。これは組織内に限った話じゃなくて、外注との関係にも通じる。
この人できるプロジェクトマネージャーだと感じる発注元の担当者は、決して敵対せずに「一緒にプロジェクトを成功させよう」というスタンスで一貫してる。だから遅延の報告もしやすく、対策を一緒になって考えてくれる。
採用が大事 全話見た今だから言えるが、高梨がいろいろ残念だというのが一番の問題。2クール目で出てくる「誰を採っても太郎よりマシ」という評価が全てな感ある。
悪い進捗の報告したときに怒る人って、「怒ると、次から怒られるのがイヤだからちゃんとスケジュールに遅れないようになるだろう」みたいな論理だと思うけど(そしてそれが「怒られるのがイヤだから嘘ついて怒られないようにしよう」になり得るのは上述の通りだけど)、そもそも問題が起きたときに反省して改善できる人であれば怒る必要は無いわけで、つまり採用がすべて。
パワーハラスメントしないと改善しないと感じてしまうような人を採用するプロセスの問題だと思う。だからと言ってパワーハラスメントする人を採用してしまってるのもまたプロセスの問題である。
まとめ  マネージャーが怒るのはスケジュール管理において悪手で責任の放棄にもつながる マネージャーと被管理者は同じ目的を目指すチームメイトという意識でいるべきで、決して敵対してはいけない 採用がすべて  </description>

            <category>プロジェクトマネジメント</category>

            <category>SHIROBAKOで学ぶプロジェクトマネジメント</category>

        </item>

        <item>
            <title>UnityのエディタとしてJetBrainsのRiderを使う</title>
            <link>http://localhost:1313/2016/09/rider_with_unity/</link>
            <pubDate>Fri, 02 Sep 2016 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2016/09/rider_with_unity/</guid>
            <description>JetBrainsのIDEって最高ですよね Mono DevelopのViモードがVimmerにとってはすごく微妙やけども使わないわけにはいかないという昨今、IdeaVimという最高のVimキーバインドプラグインを擁するJetBrainsのIDEがついにC#に対応して2016/08リリース予定でmacOS使い大歓喜にもかかわらず、未だにリリースされてないRiderをUnity開発で使う話。
ブレークポイントを張れるようにするまで。
RiderをUnityプロジェクトで使う A cross-platform C# IDE based on the IntelliJ platform and ReSharperにアクセスし、下記の手順で実行ファイルのURLがメールで送られてくるようになる。
 Get an early buildをクリックする フォームを入力する SUBSCRIBEをクリックする  ダウンロードしてインストールしておく。
Macの場合はMonoをインストールする $ brew install mono  Windowsの場合はそもそもなんかする必要があるのかわかりません。
RiderのメニューでPreferences→Builde, Execution, Deployment→Monoで/usr/local/opt/monoを設定する。
Unityプロジェクトにプラグインを追加してデバッグを有効にする  JetBrains/Unity3dRider: Unity JetBrains Rider integrationのAssets/Plugins/Editor/RiderUnityIntegration.csをUnityプロジェクトに追加する UnityのメニューでPreferences→External Toolsをクリックする Browseをクリックし、インストールされたRider.exe(Windows)/.lnkまたはrider.sh/Rider.app(Mac)を選択する UnityでPreferences→External ToolsのEditor Attachingにチェックを入れる  チェックが入ってなかった場合はUnityを再起動する   Riderでデバッグの設定  アクティビティモニタを開き、UnityのPIDを確認する RiderのメニューでRun→Edit Configurations +をクリックしてMono remoteをクリック Portに、56000 + UnityのPIDの下3桁の数字を入力してOKをクリック  例 PIDは5323なので、Mono remoteのポートは56000 + 323 = 56323を指定する。</description>

            <category>Rider</category>

            <category>Unity</category>

        </item>

        <item>
            <title>情報会議 #2 に参加しました #Johokaigi</title>
            <link>http://localhost:1313/2015/11/information_sharing2/</link>
            <pubDate>Fri, 20 Nov 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/11/information_sharing2/</guid>
            <description>24人で考えた情報共有ツールの悩み 情報会議 #2 レポート | JohoKaigi - 情報会議に行ってきた。
どんな感じやったかは上記記事見ていただいた方が早い感じあるので思ったことつらつらと書く。
esa.io  Qiita:Teamの中でesa.ioの中の人がLTするという器の大きさ。
esaは前職で使ってた（esa.io でドキュメント共有したらものすごく捗ってる | mizoguche.info）。使い始めていい感じやな～って実感して記事増えてきてカテゴリ分けどうしようくらいで転職したのでディープに使いこなすとこまではいってない感じではある。
そういえばGoogleで検索した時にesa内の検索結果を表示するChrome拡張 もざっくり書いたりしてたけど、結局転職してあんまり使う機会がなくなってしまった。
スライドにもある通り、WIPという「書きかけやけど公開しとくねー」ってノリで投稿して、誰でも編集できるから育てていこうぜ！ っていう機能が良い。投稿へのハードルがめっちゃんこ低くなってる。
この辺、Qiita:Teamは編集リクエストはできるけど誰でも編集できるわけではないので思想の違いが見える。
Qiita:Teamは個人がドキュメントを所有して共有してく、esaはチームでドキュメント所有してチームでドキュメントをつくっていく、って感じやろうか。
あと、esaの（ブラウザ上の）エディタは細かいところまで気が利いてて良い。チェックリストの箇条書きしてたら改行時に次の行もちゃんとチェックリストになってくれたり、チェックをトグルするショートカットがあったり。Kobito頑張れ。
情報共有はコミュニケーション  @hapickyさんのLTの
 情報共有のレベルが 信頼関係のレベルを 上回ることはない
情報共有≦信頼関係
 ってとこらへんで、「情報共有」の本質は「情報伝達」＝「コミュニケーション 」だということに気づかされた。
 辞典類ではまず、人間の間で行われる知覚・感情・思考の伝達、などといった簡素な定義文が掲載されている。
ただし、上記のような定義文では不十分で、一般に「コミュニケーション」というのは、情報の伝達だけが起きればが充分に成立したとは見なされておらず、人間と人間の間で、《意志の疎通》が行われたり、《心や気持ちの通い合い》が行われたり、《互いに理解し合う》ことが起きて、はじめてコミュニケーションが成立した、とされている、といった説明を補っているものもある
コミュニケーション - Wikipedia
 で、「情報共有」の正体が「コミュニケーション」であるなら、理解を深めるにあたってコミュニケーション学をあたるのが近道な気がした。
先人が突き止めたコミュニケーションの本質を学んで、現代のツールを使ってより良くしてく的な。いやすでにそういう研究もあるんでしょうけど。
自分はそこら辺の話は、マクルーハンが「メディアはメッセージである」って言ってはった、くらいの雑な知識しかないのでどっから手をつけたもんかなぁっていう。
情報共有ツールはメッセージである @papixさんがLTでちらっと言ってはった「Qiita:Teamはエンジニアにとっての福利厚生」という話がその「メディアはメッセージである」って話につながってて面白い。
 普通、メディアとは「媒体」を表すが、その時私たちはメディアによる情報伝達の内容に注目する。しかし、彼はメディアそれ自体がある種のメッセージ（情報、命令のような）を既に含んでいると主張した。例えば、同じニュース内容でもメディアが新聞か放送か週刊誌かネットかで受け止め方が違ってくる。
マーシャル・マクルーハン - Wikipedia
 「Qiita:Teamというメディアを使う」ということ自体がすでに「エンジニア御用達のMarkdownでドキュメントを手軽に残せる環境を用意してますよ！ ウチはエンジニアのこと考えてますよ！」ってメッセージを含んでるんですよね。
なので、普段からMarkdownを使ったりQiitaに慣れ親しんでるエンジニアにとってはポジティブなメッセージになるし、Markdownを知らなかったりQiitaを知らなかったりする人にとっては謎のツールを導入されてて覚えなあかんくて不安というネガティブなメッセージにもなりうるという。
みたいなことを考えてると、やっぱり情報共有についての知見を深めるならコミュニケーション学を知った方がいい気がする。
世界観の理解が必要ではないか説 懇親会で情報共有について価値を見いだせない人や情報共有したがらない人にいかに情報共有してもらうか、みたいな話をしてた。
「こういうメリットがあるからやろうYO」的なアプローチはワークショップで定量的な指標ってどんなんあるかなーとか話したので置いといて、「どうやったら情報共有させられるか」ではなく「どうしてその人は情報共有したがらないのか」をボーっと考えてた。
そういうのって、仕事へのコミットメントの差かなーと思うんすよね。「仕事楽しめてるか」と言い換えてもいいかもしれない。
仕事と趣味が一致してる人間には理解し難いことではあるんですが、一部には「仕事はツラいもの（であるべき）だ」という世界観があるように見えて、そういう世界観に生きる人たちには「情報共有して欲しい」というオファーは「ツラい仕事が増えるだけ」と感じてるのかもしれないなー、と。
前職の同僚の前職（ややこしい）の環境がまさにそんな感じやったらしくて、本人は飲み会で技術の話がしたいのに「なんで飲み会で仕事の話なんかするんよ」という雰囲気だったそうで、そういう人にとって仕事の話題のコミュニケーションってのは負担でしかないんやろなーっていう。
ってことだとすると、「仕事の話を楽しいと認識させる」か「仕事の話を楽しいと感じる人しか採用しない」という話になって、前者はしんどいからやっぱ採用大事だねという身も蓋もない思考になるのはまだ採用する段階にないスタートアップにいるからですね。
ともあれ、コミュニケーション不全はコミュニケーションでしか解決できないと思うので、情報共有に乗り気でない人と喋ってその人を理解する（そして信頼関係を築く）というのが大事な気はする。
まとめ みたいなことを考えるいい機会がもらえるJohoKaigi - 情報会議 、来週第3回が行われて、年内にあともう1回は行われるとのことなので、この手のツールや手法に興味がある方は参加すると幸せになれると思うので是非参加されたし。</description>

            <category>Qiita</category>

            <category>esa.io</category>

            <category>ポエム</category>

            <category>マネジメント</category>

            <category>コミュニケーション</category>

        </item>

        <item>
            <title>情報共有ツールお悩みNight#1に参加しました</title>
            <link>http://localhost:1313/2015/11/information_sharing/</link>
            <pubDate>Thu, 05 Nov 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/11/information_sharing/</guid>
            <description>第1回「情報共有ツールお悩みNight」を開催しました - Qiita Blogに行ってた。もう１週間も経ったか。
事前準備  準備が周到すぎる Qiita:Team、Slackが用意されて事前にコミュニケーションが図られた リハーサルやってフィードバックループまわすとかイベンターの鑑すぎて超感動した  ワークショップ  各班にファシリテーター（自称世話役）がいたので進行が非常にスムーズ ファシリテーターがいなかったらこんなにスムーズには行かなかったろうなという意味でも準備最高すぎ 問題が架空の話だったので、提示されてない条件とかを考えだしてうーんみたいなことにちょこちょこなってたんですが、リアルな話をするのも問題があるかもしれないので難しいですねーと思った  この辺はワークショップに慣れてるかどうかな感じな気もする   LT @htomine チームで考えよう   エモーショナルだけど心に刺さった 情報共有のノリが悪い人もチームの一員なので「敵」ではない 「敵」の言うことなんか聞くわけないんやから一緒に頑張っていこうというスタンスは非常に重要  @naoya@github 情報共有 失敗あるある   ウチも日報だらけのタイムラインで記事の密度が薄めやなーと思ったのでなんとかした方がいいなーと思った メンションうるさい問題みたいなのが起こる規模の環境にいたことがないので今後遭遇したら思い出したい  @tseigo 情報共有ツール導入することがある私が導入されるときの話（とくに初動面）   情報共有ツールお悩みNight #1 「情報共有ツール導入することがある私が導入されるときの話（とくに初動面）」  from Seigo Tanaka
 いろんな環境に能動てきに入れたり受動的に入れられたり 環境によってノリが違うのがツラそうな感じがある  懇親会  それぞれの環境でそれぞれに悩みを持ってる  周りが情報共有に興味がないかんじだったり 人数が増えて問題が顕在化してきてたり チームや職域によって使ってるツールが違ったり  事前にこういう問題が起きるという認識を持ててるのは今から大きくしていこうという会社にいる人間にとってはありがたい  Naoyaさんと喋った  9月に東京に引っ越してきたときに、Naoyaさんと喋って握手してもらうという目標を密かに立ててたら1ヶ月ちょいで達成してしまった  Rebuildで喋ってる人は実在した 自分の中での会いにいける（そして握手してもらえそうな）アイドルだったので超興奮して寝れなかった 「俺一般人だから」とおっしゃってましたけど、一般人がするアニメの話を聞きたいとは思わないので一般人ではないと思います  「サンフランシスコに宮川さんに会いに行くくらいの目標立てろよ」と言われてたしかにーと思ったので今後の目標はサンフランシスコで宮川さんと握手することになりました スタートアップやってると言ったら「うわーツラいねー上手くいかない時期に人間関係も悪くなるから友達失うよー」と言われたので友達失わないように頑張っていきたい 興奮して上手くしゃべれない+頭おかしいトークを多々してしまって反省しているが後悔はしていない  その後社内でやったこと 日報だらけのタイムラインをなんとかした  翌日、レトロスペクティブで下記のような問題があるなーという話になった  日報だらけで薄い感じある デザイナーが1人だけなので単一障害点になってる  解決策として、作業にとりかかりながら「タスク振り返り」という記事を書くというルールを設けた  やること、発生した問題、参考を書く  1週間くらい実施して下記のメリットを感じている  「やること」を書くのは作業にとりかかる前に行われるため、タスクの見通しが良くなり段取り力が養成される 問題点と解決策と参考URLが共有されるので知識が属人化しなくて良い  日報は日報で書く  その際、今日の「タスク振り返り」へのリンクを張る どう考えてもリンク張るのめんどくさいので自分が今日Qiita:Teamへ投稿した記事へのリンクをコピーするChrome拡張つくった  とりあえず4人のチームだから上手くいってるのかもしれないけど、もし他の環境に放り込まれてもQiita:Team入れてこのルール導入したいなーとは思う  第3回もあるよ 情報共有ツールお悩みNight #3 | Peatix</description>

            <category>Qiita</category>

            <category>ポエム</category>

            <category>マネジメント</category>

        </item>

        <item>
            <title>ソフトウェア開発の会社で働き終わって思ったこと</title>
            <link>http://localhost:1313/2015/08/quit/</link>
            <pubDate>Wed, 19 Aug 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/08/quit/</guid>
            <description>8月7日付けで退職した。
アルバイトの子から「退職エントリ待ってます」的なフリがあったし、思うことを書いとく。
ただの自戒を込めた反省なので、他人をディスる意図は全くない。
ビジネスモデル 人月商売について  このシステムは作るのに何時間かかるからウン万円という人月単価×工数=請求金額のビジネスモデルを人月商売とする 作業効率が上がって作るのに時間がかからなくなればなるほど儲からない 作業効率は技術力と比例する 低い技術力でなんとか実現することで稼げる 経営レイヤーでは（短期的に）見ると技術力を上げないインセンティブが働く  単価はできるだけ低く抑えて利益を上げたくなる 単価を上げる（技術力の高い人を雇う）と回収できるかわからないリスクがある  この機能はこれくらいの価値があるという観点から請求金額が導けると顧客も開発会社もハッピーになれそうな気がしてる  実装に時間がかかろうがかからなかろうが、機能の持つ価値が請求金額に反映される世界観 おそらく顧客は納得しにくいから実現が難しい   他社との差別化  技術力が他社と変わらないか他社よりも低いと価格で競争するしかなくなる 景気が悪くなって需要がなくなったら死ねそう 「技術」を拠り所にしないとサステナビリティはなさそう  ここで言う「技術」は、価値あるソフトウェアを実現するための非常に広い範囲を含んだ技術（高い価値をもつデザイン・ユーザービリティetc.を実現する技術） 価値が（安い）価格にしかないのであればリピーターがつきにくいし、リピーターも価格で劣ればリピートしなくなりそう   経験がある人は（少なくともすぐには）変われないし変わらない 勉強会  デザインパターン・ソフトウェアテストの社内勉強会(読書会)を行った 業務に活かされることは（少なくとも1年の期間では）ないように見えた 心の底から本当に必要だと思わなければ技術は習得されない  コードレビュー  Bitbucketでプルリクエストベースのコードレビューを行った レビューで何回もリーダブルコードに書いてあるような同じことを指摘した 心の底から本当に必要だと思わなければコードの書き方は変わらない  つまり  技術に対する価値観は年齢を重ねると変わらない 知識をアップデートする必要性を感じないまま年齢を重ねたらヤバい  必要性を感じる事自体がある意味技術かもしれない（哲学）    レベル3を超えることなく、10年過ぎてしまった人の意識を変えさせるのは、私自身の経験から、私はほとんど不可能だと思っています。 ソフトウェアエンジニアの成長カーブ（再掲載）：柴田 芳樹 (Yoshiki Shibata)：So-netブログ
  年齢を重ねた人の考え方を変えれるなどというのはおこがましい考え  年齢を重ねる前に洗脳するしかない  採用の段階で選別するのが一番効率が良いという身も蓋もない結論に至らざるをえない  コミュニケーション  コミュニケーションは伝える側が自分の考えを伝える責任と、受け取る側が伝える側の考えを受け取る責任の2つがあり、どちらが欠けても目的を果たせない 受け取る責任は、伝える側に質問する能力を持たないと果たすことができない  日常会話では、言葉の定義に揺れや曖昧さがある場合が多い ex.</description>

            <category>コードレビュー</category>

            <category>コミュニケーション</category>

            <category>マネジメント</category>

        </item>

        <item>
            <title>esa.io でドキュメント共有したらものすごく捗ってる</title>
            <link>http://localhost:1313/2015/04/esa/</link>
            <pubDate>Wed, 22 Apr 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/04/esa/</guid>
            <description> esa.io を会社で導入した。
ドキュメント管理が課題だとさんざん言いつつも放置してきており、社内に知識が溜まっていかないという残念な組織づくりが行われていた。
これはいかんと Redmine の Wiki をつかってみたりもしたけど、しっくりこない感。Redmine はプロジェクトに紐づくのでナレッジベースにするような運用できない。
Qiita:team は高いしなーという世間話から、そういえば esa.io ってあったよなー最初は無料やしいっぺん使ってみようかーとなって使ってみたら便利な雰囲気を感じ取って、今に至る。
pros  気軽に投稿できる 検索できる Markdown で書ける デザインがいい感じ(ズルいデザインテクニックの人らしい) 公式 Twitter アカウントが割と esa に関するツイートを拾ってくれるしメンションしたらすぐに返してくれる  cons  APIがない カテゴリ・タグの複数記事の一括編集など、ドキュメントの検索性・閲覧性を高める機能が貧弱というか見当たらない  印象  とりあえず日報から始めるパターンは鉄板  書いてると釣られてみんな書くっていうのは割とホンマっぽい  カテゴリ・タグの一括編集がしにくいのはかなり致命的な感じはする  ただ、記事APIは早い段階で提供したいということなので、カテゴリ・タグの複数一括編集はAPIができたら叩けば可能な期待はある   @mizoguche 申し訳ありません、現在APIの準備中で、まだ公開に至っておりません。検索APIは早い段階で記事APIと合わせて提供したいと考えています
&amp;mdash; esa_io (@esa_io) April 10, 2015
 検索APIもつくってるそうなので、Google の検索結果の横に esa の検索結果も表示するみたいなブラウザ拡張書いて捗る未来が見える  まとめ  情報をストックするツールとして現状不満もあるけど、将来的には大丈夫な感じになるとは思われる ナレッジベースを探してて Qiita はちょっとお値段張るなぁと思う層には、esa.io がフィットしてる  </description>

            <category>チーム開発</category>

            <category>esa.io</category>

            <category>ドキュメント</category>

        </item>

        <item>
            <title>Terraform で AWS をテラフォーミングしたときに困ったこと</title>
            <link>http://localhost:1313/2015/02/terraform_aws/</link>
            <pubDate>Thu, 05 Feb 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/02/terraform_aws/</guid>
            <description>AWS の設定するとき、いちいちコンソールからいじってた情弱だったが、Terraform を使い始めて文明開化の音がした。
けど、やっぱり初めて使うものですからわからんこと／困ったことがあった。
.gitignore にいれるべきではないもの terraform gitignore - Google 検索で上の方に出てくる .gitignore はアプリケーション開発用の .gitignore なのでこれは使ってはいけない。
これを使うと .tfstate ファイルが無視されてしまい、AWS の状態が保持できなくなる。
what&amp;rsquo;s a good git-ignore policy for a terraform repository? - Google グループ
 We recommend checking in both your plan files (*.tf) and your .tfstate files. This will allow others to modify the infrastructure. Without state, existing infrastructure won&amp;rsquo;t be found.
 terraform apply するたびに状態をコミットするべき、ということみたい。
データベースのパスワードをコミットするかは要検討なものの、 それ以外に.gitignore に追加するべきファイルはなさそう。
デフォルト設定では RDS で日本語が使えない デフォルトのパラメータグループだとちょこちょこ latin1 になる設定があるっぽい。これを UTF-8 に設定するパラメータグループを作る。</description>

            <category>AWS</category>

            <category>Terraform</category>

        </item>

        <item>
            <title>オープン・クローズドの原則 - アジャイルソフトウェア開発の奥義一人読書会(2)</title>
            <link>http://localhost:1313/2015/01/agile-development-reading-2/</link>
            <pubDate>Sat, 31 Jan 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/01/agile-development-reading-2/</guid>
            <description>アジャイルソフトウェア開発の奥義 第2版 オブジェクト指向開発の神髄と匠の技 著者: ロバート・C・マーチン,瀬谷 啓介 発売日: 2008-07-01   オープン・クローズドの原則（OCP: Open-Closed Principle） オープン  モジュールの振る舞いを拡張できる
 仕様変更に対して、モジュールに新たな振る舞いを追加することで対処できるようにすること。
クローズド  モジュールの振る舞いを拡張してもソースコードやバイナリが影響を受けない
 ようするに 拡張がしやすくて、拡張しても修正箇所はできるだけ少なくなるような設計にするべき、という話。
「抽象」に依存する  モジュールをある固定した「抽象」に従属させておけば、修正に対してコードを閉じることができるのだ。なぜなら、「抽象」を使えば、コードを修正しなくても、その「抽象」の派生クラスを新たに追加するだけでモジュールの振る舞いを拡張できるからである。
 修正するたびにバグを埋め込む機会を得ることになるので修正箇所はできるだけ少なくするべき。それを実現する方法が、抽象に対するプログラミング。
public class Client { private Server server; public Client(){ server = new Server(); } public void someMethod(String param){ String str = server.process(param); // ... } } public class Server { public String process(String param){ // ... } // 他の必要なメソッド }  このように、具体的なクラスに依存している Client クラスは別の Server オブジェクトを利用したくなった場合はClient クラスを修正する必要がある。</description>

            <category>アジャイルソフトウェア開発の奥義</category>

            <category>オブジェクト指向</category>

            <category>本</category>

            <category>設計</category>

        </item>

        <item>
            <title>Unity の Android プラグインつくった</title>
            <link>http://localhost:1313/2015/01/unity_android_plugin/</link>
            <pubDate>Sun, 25 Jan 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/01/unity_android_plugin/</guid>
            <description>iOS / Android でツイートしたらコールバック貰える Unity のアセット探してたけどなかったのでつくった。
mizoguche/TweetSharer
デザインパターン厨としてはここで使わないとどこで使うって感じだったので Bridge パターンを使ってみた。
やっぱり偉い人たちが言うだけあって、かなりコードが見やすくなって気持ちいい。
Bridge パターン使わなかったら #if UNITY_ANDROID と #elsif UNITY_IOS による精神汚染が激しくなってたと思われる。
iOS 対応は anchan828/social-connector を参考にした。
Android 対応 めんどくさい部分 Android でただただツイートを他アプリに任せるなら暗黙的インテント使えばいいだけの話やけど、今回はツイートしたのかキャンセルしたのかを知りたいので startActivityForResult で Activity 呼んで onActivityResult で受け取って、みたいなめんどくささがある。
Android Studio で jar をつくるよ Android Studio のプロジェクトをつくる リポジトリのルートが Unity プロジェクトのルートで、その直下に Androi Studio プロジェクトのディレクトリを置いた。
特にこれが良さそうという根拠はなかったけど、今のところこれが良さそう。理由は後述。
ビルドパスに Unity の classes.jar を追加する Unity の Java クラスを参照したいので、 Unity の classes.jar を ビルドパスに追加する。
Unity で Android のネイティブプラグインの実装 - ローカル通知を送信する - nirasan&amp;rsquo;s tech blog</description>

            <category>Android</category>

            <category>Unity</category>

        </item>

        <item>
            <title>EC2 のインスタンスを Itamae でプロビジョニングした</title>
            <link>http://localhost:1313/2015/01/provisioning_with_itamae/</link>
            <pubDate>Sat, 10 Jan 2015 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2015/01/provisioning_with_itamae/</guid>
            <description>今までプロビジョニングツールには Chef 使ってた（使いこなせてなかった）。小規模なサーバーしか扱わないのでレシピ組み立てるのに HP がかなり必要だった。
軽量シンプル Chef ともっぱらの評判の、itamae-kitchen/itamae を使ってみた。
レシピの書き方 パッケージのインストール package &#39;git gcc-c++ patch readline readline-devel zlib&#39; do action :install end  例によって、これで入る。
ファイルのアップロード remote_file &#39;/home/ec2-user/.bash_profile&#39; do source &#39;bash_profile&#39; end  例によってこれで上がる。
Resources · itamae-kitchen/itamae Wiki にドキュメントが上がってないので、itamae/file.rb at master · itamae-kitchen/itamae と itamae/remote_file.rb at master · itamae-kitchen/itamae を見ながら使い方を確認した。
remote_file &#39;/etc/sudoers&#39; do source &#39;sudoers&#39; mode 440 owner root group root end  こうすれば mode、owner、groupを設定できるので /etc/sudoers を上書きしても大丈夫（1回モードもオーナーもグループも設定せずに上書きして詰んだ）。
レシピのインクルード rbenv のインストールに k0kubun/itamae-plugin-recipe-rbenv をつかった。</description>

            <category>Itamae</category>

            <category>サーバー</category>

        </item>

        <item>
            <title>Unity で APK が 50MB 超えたので OBB Downloader 使う</title>
            <link>http://localhost:1313/2014/12/unity_android_50mb_over/</link>
            <pubDate>Thu, 25 Dec 2014 00:00:00 +0000</pubDate>

            <guid>http://localhost:1313/2014/12/unity_android_50mb_over/</guid>
            <description>APK が 50MB を超えた APKが50MBを超えてしまうとGoogle Play Storeにアップロードできない。
Unity - Manual: Support for Split Application Binary (.OBB) にある通り、Unityにはビルド時にAPKを分割する機能がある。
注意点  Google Play Storeの仕様では、基本的にAPKと分割されたOBBファイルはインストール時に同時にダウンロードされて所定の位置に保存されるが、条件によってはOBBがダウンロードされずAPKのみダウンロードされる場合がある。 よって、アプリ起動時にOBBファイルが所定の位置にあるか確認する必要がある 参考: APK Expansion Files | Android Developers  Unity ではどうするの Unity には Google Play OBB Downloader by Unity Technologies &amp;ndash; Unity Asset Store というアセットがあるのでこれを利用するとよい。
処理の手順は下記の通り。
 最初に OBB がダウンロードされているか確認するシーンを表示する OBB がダウンロードされていればタイトルなり最初に表示したいシーンに遷移させる なければ OBB のダウンロードを開始するように実装する。  Tutorial Unity 4 apk splitting into OBB for google play にあるコードが最初のシーンのコードとして参考になる。</description>

            <category>Android</category>

            <category>Unity</category>

        </item>

    </channel>
</rss>
	`
