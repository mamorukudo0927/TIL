# ターミナル操作の基礎学習
このドキュメントでは、ターミナル操作の基礎について学習した内容を記録する。
ターミナルとは、**キーボードから「コマンド」という命令文を入力し
コンピュータに命令をするアプリケーションのこと**を指す。
ここでは、いくつかのコマンド操作を例に学習を進める。

# 日常で利用するターミナル
Windowsであれば「コマンド・プロンプト」「Pwower Shell」、Macであれば「ターミナル」アプリケーションを利用してコマンド操作が可能。
例えば本ファイルは以下のPwowerShellコマンドで作成できる。

```
# 現在のパスにディレクトリ作成を行う。
mkdir Rquier

# cd 以降に指定したパスに移動する。
cd C:\your\path\dir\

# new-item：ファイルを新規作成する。命令
new-item BasicTerminalUsage.md
```

# プログラミング言語関連のターミナル
例えば、
- Node.jsのパッケージマネージャ**npm**や
- Pythonのパッケージマネージャ**pip**

等はコマンド操作によりライブラリの取得・インストール・依存関係の整理を行うことができる。

これらはあらかじめ予約語と予約語に対応する処理を定義されており、利用者は処理内容を気にすることなく指定されたコマンドを入力するのみで結果を得ることができる。

```
# PythonのWebFramework 「Flask」のインストール
pip install Flask

# JavaScriptのViewFramework「React.js」のインストール
npm install react
```

コマンド操作はエンジニアではないユーザの利用は難しい場合があるが、エンジニアであれば避けては通れない。

現代ではブラウザを介してGUIベースの操作、PCにGUIアプリケーションをインストールして操作といった、コマンド操作を隠ぺいすることも多いが、

コンピュータができた当時はコマンド操作が一般的であり、
コマンド操作を行うほうがPCとしては自然。

# コマンド操作を受け付けるアプリケーションを実装する。
ここではPythonを用いて、コマンド操作を受付け、結果を入力内容によって可変させる処理を実装する。

```python basicTerminalUsage.py
import click

@click.command()
@click.option('--name', required=True, type=str)
def cmdline(name):
    click.echo('Hello {name}!'.format(name=name))

if __name__ =='__main__':
    cmdline()
```

実行結果は以下の通り

```
python basicTerminalUsage.py --name 'mkudo'
Hello mkudo!

python basicTerminalUsage.py
Try 'basicTerminalUsage.py --help' for help.
Error: Missing option '--name'.

python basicTerminalUsage.py --help
Options:
  --name TEXT  [required]
  --help       Show this message and exit.
```

このように、

- あらかじめ指定した関数を
- あらかじめ指定した呼び出しで

実行していることがわかる。
--helpに関しては利用しているライブラリ「Click」の中で定義されている。

