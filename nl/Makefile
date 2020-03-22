NPM_BIN = ../node_modules/.bin
SIMPLIFY ?= 100

%_raw: %.zip
	unzip $< -d $(basename $<)_raw || true

# 小学校区 2016
A27-16_%_GML.zip:
	curl -o $(notdir $@) -LO http://nlftp.mlit.go.jp/ksj/gml/data/A27/A27-16/$(notdir $@)

A27-16_%_GML_Surface.geojson: A27-16_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/shape/A27-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

A27-16_%_GML_Point.geojson: A27-16_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/shape/A27P-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

# 中学校区 2016
A32-16_%_GML.zip:
	curl -o $(notdir $@) -LO http://nlftp.mlit.go.jp/ksj/gml/data/A32/A32-16/$(notdir $@)

A32-16_%_GML_Surface.geojson: A32-16_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/shape/A32-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

A32-16_%_GML_Point.geojson: A32-16_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/shape/A32P-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

# 行政区域 2019
N03-190101_%_GML.zip:
	curl -o $(notdir $@) -LO http://nlftp.mlit.go.jp/ksj/gml/data/N03/N03-2019/$(notdir $@)

N03-190101_%_GML.geojson: N03-190101_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

# 鉄道 2018
N02-18_GML.zip:
	curl -o $(notdir $@) -LO http://nlftp.mlit.go.jp/ksj/gml/data/N02/N02-18/$(notdir $@)

N02-18_GML_Station.geojson: N02-18_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/*_Station.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

N02-18_GML_RailroadSection.geojson: N02-18_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/*_RailroadSection.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@
