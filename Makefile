
DATAS = datas
ARTIFACTS = artifacts

build-geojson:
	../node_modules/.bin/mapshaper shape/A27-16_13.shp encoding=shift-jis -o out.geojson

# 小学校区 http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-A27-v2_1.html
${DATAS}/A27-16_13_GML.zip:
	TARGET='http://nlftp.mlit.go.jp/ksj/gml/data/A27/A27-16/A27-16_13_GML.zip' && \
	curl -o ${DATAS}/`basename $${TARGET}` -LO $${TARGET}

${ARTIFACTS}/A27-16_13.geojson: ${DATAS}/A27-16_13_GML.zip
	unzip ${DATAS}/A27-16_13_GML.zip
	./node_modules/.bin/mapshaper ${DATAS}/shape/A27-16_13.shp encoding=shift-jis -o out.geojson
