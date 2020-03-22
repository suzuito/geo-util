# 地理データユーティリティ

- node 12.12.0 or later

## Set up

```sh
npm install
```

## [国土数値情報のデータ](http://nlftp.mlit.go.jp/ksj/index.html)からGeoJSONを生成する

国土数値情報のデータからGeoJSONを生成するプログラム。このプログラムはnlディレクトリの配下にある。

```sh
# プログラムを利用するためには、nlディレクトリに移動する。
cd nl
```

- 下のサンプルコードの`${pref}`は都道府県コード。`01-47`。
- geojsonのシンプル度（データ量削減）
  - 環境変数`SIMPLIFY`に`1-100`を設定し、makeを実行する。例えば`SIMPLIFY`を`10`にした場合、90%のデータを削減した.geojsonを生成する。何も設定しない場合、`SIMPLIFY`は`100`となる。

### [小学校区データ 2016年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-A27-v2_1.html)

```sh
# 学区を表す面データ
make A27-16_${pref}_GML_Surface.geojson
# 学校の場所を表す点データ
make A27-16_${pref}_GML_Point.geojson
```

### [中学校区データ 2016年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-A32-v2_0.html)

```sh
# 学区を表す面データ
make A32-16_${pref}_GML_Surface.geojson
# 学校の場所を表す点データ
make A32-16_${pref}_GML_Point.geojson
```

### [行政区域データ 2019年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03-v2_3.html)

```sh
# 行政区を表す面データ
make N03-190101_${pref}_GML.geojson
```

### [鉄道データ 2019年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N02-v2_3.html)

```sh
# 鉄道の駅を表す線分データ
make N02-18_GML_Station.geojson
# 鉄道の路線を表す線分データ
make N02-18_GML_RailroadSection.geojson
```
