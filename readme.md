# 步骤
```
#1. 构建镜像 Dockerfile
docker build -t my/tesseract-ocr .

#2. 修改镜像名
vi docker-compose.yml
    image: my/tesseract-ocr
    
#3. 运行
docker-compose up -d
```
# main.go
对外提供一个接口`/ocr?url=http://xx.com/aa.jpg`

# 测试
浏览器访问`http://localhost:8000/ocr?url=http://xx.com/aa.jpg`
```
{
	"code": 200,
	"msg": "success",
	"data": {
		"tips": "Warning: Invalid resolution 0 dpi. Using 70 instead.\nEstimating resolution as 197\n",
		"content": "PHOTOGRAPH\n产品实拍\n\n由于拍摄光线以及显示器等因素影响，可能会导致图片与实物颜色\n有细微信差，最终颜色以实物为准。\n\u000c",
		"url": "http://xx.com/aa.jpg"
	}
}

```
后来发现，在没优化的前提下，chineseocr_lite的效果会更好，详见
http://www.cuiwei.net/p/1052444754
