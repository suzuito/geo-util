# 地理データユーティリティ

- node 12.12.0 or later

## Set up

```sh
npm install
```

## 国土数値情報のデータからGeoJSONを生成する

http://nlftp.mlit.go.jp/ksj/index.html

### [小学校区データ 2016年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-A27-v2_1.html)

```sh
cd nl
for pref in `seq -w 1 47`;
do
    make A27-16_${pref}_GML.geojson
done
```

### [中学校区データ 2016年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-A32-v2_0.html)

```sh
cd nl
for pref in `seq -w 1 47`;
do
    make A32-16_${pref}_GML.geojson
done
```

### [行政区域データ 2019年](http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03-v2_3.html)

```sh
cd nl
for pref in `seq -w 1 47`;
do
    make N03-190101_${pref}_GML.geojson
done
```
