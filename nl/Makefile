NPM_BIN = ../node_modules/.bin
SIMPLIFY ?= 100
TMP_DIR = ./tmp
RESULT_DIR = ./result

%_raw: %.zip
	unzip $< -d $(basename $<)_raw || true

$(TMP_DIR):
	mkdir $(TMP_DIR)

$(RESULT_DIR):
	mkdir $(RESULT_DIR)

clean:
	rm -rf $(TMP_DIR) $(RESULT_DIR)

# 行政区 2019

$(TMP_DIR)/N03-190101_%_GML.zip: $(TMP_DIR)
	curl -o $@ -LO https://nlftp.mlit.go.jp/ksj/gml/data/N03/N03-2019/$(notdir $@)

$(RESULT_DIR)/N03-190101_%_GML.geojson: $(TMP_DIR)/N03-190101_%_GML_raw $(RESULT_DIR)
	cp $</*.geojson $@

N03-190101:
	for i in `seq -w 1 47`; do make $(RESULT_DIR)/N03-190101_$${i}_GML.geojson; done

# 小学校区 2016
$(TMP_DIR)/A27-16_%_GML.zip:
	curl -o $@ -LO https://nlftp.mlit.go.jp/ksj/gml/data/A27/A27-16/$(notdir $@)

$(RESULT_DIR)/A27-16_%_GML_Surface.geojson: $(TMP_DIR)/A27-16_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/shape/A27-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

$(RESULT_DIR)/A27-16_%_GML_Point.geojson: $(TMP_DIR)/A27-16_%_GML_raw
	$(NPM_BIN)/mapshaper $(basename $<)/shape/A27P-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

A27-16:
	for i in `seq -w 1 47`; do make $(RESULT_DIR)/A27-16_$${i}_GML_Point.geojson; done
	for i in `seq -w 1 47`; do make $(RESULT_DIR)/A27-16_$${i}_GML_Surface.geojson; done

# 中学校区 2016
$(TMP_DIR)/A32-16_%_GML.zip:
	curl -o $@ -LO https://nlftp.mlit.go.jp/ksj/gml/data/A32/A32-16/$(notdir $@)

$(RESULT_DIR)/A32-16_%_GML_Surface.geojson: $(TMP_DIR)/A32-16_%_GML_raw
	$(NPM_BIN)/mapshaper $</shape/A32-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

$(RESULT_DIR)/A32-16_%_GML_Point.geojson: $(TMP_DIR)/A32-16_%_GML_raw
	$(NPM_BIN)/mapshaper $</shape/A32P-*.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

A32-16:
	for i in `seq -w 1 47`; do make $(RESULT_DIR)/A32-16_$${i}_GML_Point.geojson; done
	for i in `seq -w 1 47`; do make $(RESULT_DIR)/A32-16_$${i}_GML_Surface.geojson; done

# 鉄道 2018
$(TMP_DIR)/N02-18_GML.zip:
	curl -o $@ -LO https://nlftp.mlit.go.jp/ksj/gml/data/N02/N02-18/$(notdir $@)

$(RESULT_DIR)/N02-18_GML_Station.geojson: $(TMP_DIR)/N02-18_GML_raw
	cp $</N02-18_Station.geojson $@

$(RESULT_DIR)/N02-18_GML_RailroadSection.geojson: $(TMP_DIR)/N02-18_GML_raw
	cp $</N02-18_RailroadSection.geojson $@

N02-18:
	make $(RESULT_DIR)/N02-18_GML_Station.geojson
	make $(RESULT_DIR)/N02-18_GML_RailroadSection.geojson

# 河川 年代バラバラ
# $(TMP_DIR)/W05-08_%_GML.zip:
# 	curl -o $@ https://nlftp.mlit.go.jp/ksj/gml/data/W05/W05-08/$(notdir $@)
# 
# W05-08_%_GML_Stream.geojson: W05-08_%_GML_raw
# 	$(NPM_BIN)/mapshaper $(basename $<)/*_Stream.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@
# 
# W05-08_%_GML_RiverNode.geojson: W05-08_%_GML_raw
# 	$(NPM_BIN)/mapshaper $(basename $<)/*_RiverNode.shp encoding=shift-jis -simplify $(SIMPLIFY)% -o $@

# 高速道路 2018
$(TMP_DIR)/N06-18_GML.zip:
	curl -o $@ https://nlftp.mlit.go.jp/ksj/gml/data/N06/N06-18/$(notdir $@)

$(RESULT_DIR)/N06-18_HighwaySection.geojson: $(TMP_DIR)/N06-18_GML_raw
	cp $</N06-18_HighwaySection.geojson $@

$(RESULT_DIR)/N06-18_Joint.geojson: $(TMP_DIR)/N06-18_GML_raw
	cp $</N06-18_Joint.geojson $@
