<template>
  <div class="hello">
    <p>github repository is 
      <a href="https://github.com/po3rin/vue-golang-fileserver" target="_blank" rel="noopener">po3rin/vue-golang-fileserver</a>.
    </p>
    <vue-dropzone ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"
      v-on:vdropzone-removed-file="removeEvent"
      v-on:vdropzone-sending="sendingEvent"
    ></vue-dropzone>
  </div>
</template>

<script>
import vue2Dropzone from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
export default {
  name: 'HelloWorld',
  components: {
    vueDropzone: vue2Dropzone
  },
  data: function () {
    return {
      dropzoneOptions: {
        url: 'http://localhost:5000/upload',
        method: 'post',
        addRemoveLinks: 'true'
        // autoQueue: false,
        // autoProcessQueue: false
        // thumbnailWidth: 150,
        // maxFilesize: 0,
        // maxFiles: 1,
        // headers: { "My-Awesome-Header": "header value" }
      }
    }
  },
  mounted () {
    this.axios.get('http://localhost:5000/list').then(res => {
      res.data.forEach(res => {
        let filename = res.path.replace('http://localhost:5000/assets/', '')
        let id = filename.replace('.png', '')
        var file = {size: res.size, name: filename, upload: {uuid: id}}
        this.$refs.myVueDropzone.manuallyAddFile(file, res.path)
      })
    }).catch(err => {
      console.log(err)
    })
  },
  methods: {
    removeEvent: function (file, error, xhr) {
      this.axios.delete(`http://localhost:5000/delete/${file.upload.uuid}`).then(response => {
        console.log('success upload')
      }).catch(err => {
        console.log(err)
      })
    },
    sendingEvent (file, xhr, formData) {
      formData.append('uuid', file.upload.uuid)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>