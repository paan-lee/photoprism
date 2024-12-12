<template>
  <div class="p-page p-page-photo-detail">
    <v-toolbar flat color="secondary" :dense="$vuetify.breakpoint.smAndDown">
      <v-toolbar-title>
        <translate>Photo Detail</translate>
      </v-toolbar-title>

      <v-spacer></v-spacer>
    </v-toolbar>
    <v-container fluid class="pa-4 text-selectable">
      <div class="image-detail">
        <div v-if="loading">Loading...</div>
        <div v-else-if="error" class="error">{{ error }}</div>
        <div v-else>
          <h2>{{ photo.Title || 'Untitled' }}</h2>
          <button @click.exact="copyText(photo.UID, 'Url')" title="Photo ID">
            <i>vpn_key</i>
            {{ photo.UID }}
          </button>
          <br />
          <button @click.exact="copyText(photo.Path, 'Path')" title="Image Path">
            <i>folder</i>
            {{ photo.Path }}
          </button>
          <br />
          <br />
          <div class="image-container">
            <img 
              :key="photo.Files[0].Hash"
              :src="`/api/v1/t/${photo.Files[0].Hash}/${previewToken}/fit_720`" 
              :alt="photo.Title || 'Image'"
              :title="photo.Files[0].Hash || 'Image'" />
          </div>
          <br />
        </div>
      </div>
    </v-container>

    <footer></footer>
  </div>
</template>

<script>
import useClipboard from 'vue-clipboard3'

export default {
  name: "PhotoDetail",
  data() {
    return {
      authToken: null,
      previewToken: null,
      photo: null,
      loading: true,
      error: null,
    };
  },
  setup() {
    const { toClipboard } = useClipboard(); // Import `toClipboard` from vue-clipboard3
    return { toClipboard };
  },
  created() { 
    this.authToken = localStorage.getItem('authToken');
    this.previewToken = __CONFIG__.previewToken;
    const uid = this.$route.query.uid; // Access route query parameters
    if (uid) {
      this.fetchPhotoDetails(uid);
    } else {
      this.error = "No UID provided.";
      this.loading = false;
    }
  },
  methods: {
    async fetchPhotoDetails(uid) {
      try {
        const response = await fetch(`/api/v1/photos/${uid}`, {
          method: "GET",
          headers: {
            "Authorization": `Bearer ${this.authToken}`,
            "Content-Type": "application/json"
          }
        });

        if (!response.ok) {
          throw new Error(`Error fetching image: ${response.statusText}`);
        }
        this.photo = await response.json();
      } catch (err) {
        this.error = err.message;
      } finally {
        this.loading = false;
      }
    },
    async copyText(text, condition) {
      var newText;
      switch(condition) {
        case 'Url':
          var uid = text;
          newText = `http://10.0.0.239:2342/library/photo-detail?uid=` + uid;
          break;
        default:
          newText = text;
      }

      try {
        await this.toClipboard(newText);
        alert("Copied to clipboard: " + newText);
      } catch (e) {
        console.error("Failed to copy text:", e);
      }
    },
  },
};
</script>