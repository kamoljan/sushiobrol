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
    [FIDs]
    {
    "status": "OK",
    "data": {
        "image": [
            {
                "field": "xlarge_view",
                "value": "0001-webp-xlarge-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-640-871"
            },
            {
                "field": "xlarge_list",
                "value": "0001-webp-xlarge-list-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-320-435"
            },
            {
                "field": "large_view",
                "value": "0001-webp-large-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-480-653"
            },
            {
                "field": "large_list",
                "value": "0001-webp-large-list-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-240-327"
            },
            {
                "field": "medium_view",
                "value": "0001-webp-medium-view-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-320-435"
            },
            {
                "field": "medium_list",
                "value": "0001-webp-medium-list-da39a3ee5e6b4b0d3255bfef95601890afd80709-EBEBEB-160-218"
            }
        ]
    }
    }

### GET /[FID] HTTP/1.1

	Content-Type:image/jpeg
	HTTP/1.1 200 OK
	[file content]


### PARSE
{
	"xlarge_view": "0001-webp-xlarge-view-04...8b-ACA0AC-640-543"
	"xlarge_list": "0001-webp-xlarge-list-04...8b-ACA0AC-320-284"
	"large_view": "0001-webp-large-view-04...8b-ACA0AC-480-320"
	"large_list": "0001-webp-large-list-04...8b-ACA0AC-240-190"
	"medium_view": "0001-webp-medium-view-04...8b-ACA0AC-320-256"
	"medium_list": "0001-webp-medium-list-04...8b-ACA0AC-160-102"
}

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

