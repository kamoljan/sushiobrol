# Image specification

    | Density                 | Thumbnail | Ad view |
    |-------------------------|-----------|---------|
    | xlarge (xhdpi): 640x960 | 320       | 640     |
    | large (hdpi): 480x800   | 240       | 480     |
    | medium (mdpi): 320x480  | 160       | 320     |


# API

### PUT / HTTP/1.1
	
    Content-Length: 65534
    [file content]
    
    HTTP/1.1 200 OK
    [IMAGE META]
    {
        "image": {
        	"machine": "0001",
            "hash": "8787bec619ff019fd17fe02599a384d580bf6779",
            "color": "ACA0AC",
            "width": 840,
            "height": 756,
           	"density":[
           	    {
           	    	"name": "xlarge",
           	    	"value": [
           	    	    {"ui": "view", "value": {"width": 640, "height": 543}},
           	    	    {"ui": "list", "value": {"width": 320, "height": 284}}
           	    	]
           	    },
           	    {
           	        "name": "large",
           	    	"value": [
           	    	    {"ui": "view", "value": {"width": 480, "height": 320}},
           	    	    {"ui": "list", "value": {"width": 240, "height": 190}}
           	    	]
           	    },
           	    {
           	    	"name": "medium",
           	    	"value": [
           	    	    {"ui": "view", "value": {"width": 320, "height": 256}},
           	    	    {"ui": "list", "value": {"width": 160, "height": 102}}
           	    	]
           	    }
            ]
        }
    }

### GET /[FID] HTTP/1.1

	Content-Type:image/jpeg
	HTTP/1.1 200 OK
	[file content]


### FID

0001/webp/origin/04...8b-ACA0AC-640-543
0001/webp/xlarge/view/04...8b-ACA0AC-640-543
0001/webp/xlarge/list/04...8b-ACA0AC-320-284
0001/webp/large/view/04...8b-ACA0AC-480-320
0001/webp/large/list/04...8b-ACA0AC-240-190
0001/webp/medium/view/04...8b-ACA0AC-320-256
0001/webp/medium/list/04...8b-ACA0AC-160-102

### PARSE

	image_meta:[IMAGE META]
	<!-- generate it in Parse
    xlarge_view:0001040d...8b-ACA0AC-640-543
    xlarge_list:0001040d...8b-ACA0AC-320-284
    large_view:0002040d...8b-ACA0AC-480-320
    large_list:0001040d...8b-ACA0AC-240-190
    medium_view:0001040d...8b-ACA0AC-320-256
    medium_list:0002040d...8b-ACA0AC-160-102
    -->

### DIR
     
	0001/
	    webp/
	        origin/04/0d/0001-04...8b-ACA0AC-640-543
	        xlarge/view/04/0d/0001-04...8b-ACA0AC-640-543
	        xlarge/list/04/0d/0001-04...8b-ACA0AC-320-284
	        large/view/04/0d/0001-04...8b-ACA0AC-480-320
	        large/list/04/0d/0001-04...8b-ACA0AC-240-190
	        medium/view/04/0d/0001-04...8b-ACA0AC-320-256
	        medium/list/04/0d/0001-04...8b-ACA0AC-160-102
        jpeg/

### TEST

	curl -XPUT http://localhost:9090/ -H "Content-type: image/jpeg" --data-binary @gopher.png
	curl -v -XPUT -include --form key1=value1 --form upload=@gopher.png http://localhost:9090/

