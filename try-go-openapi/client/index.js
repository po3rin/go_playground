var openapi = require('api')
var api = new openapi.DefaultApi()
var postRequest = openapi.PostRequest.constructFromObject(
    {
        post: {
            title: 'hello',
            content: 'world'
        }
    }
)

api.createPost(
    postRequest,
    function(err, data, res){
        if (err) {
            console.error(err)
        } else {
            console.log('res: ' + res.text)
        }
    }
)
