
cd gut/ui
npm install -g @vue/cli
//for rapid prototyping
npm install -g @vue/cli-service-global
//vue create --no-git gut
//vue serve


cd web

//TODO add version to install for reproducible builds
cat <<EOF > vue-preset.json
 {
      "useConfigFiles": true,
      "plugins": {
        "@vue/cli-plugin-babel": {},
        "@vue/cli-plugin-eslint": {
          "config": "base",
          "lintOn": [
            "save"
          ]
        },
        "@vue/cli-plugin-unit-jest": {},
        "@vue/cli-plugin-e2e-cypress": {}
      },
      "router": true,
      "routerHistoryMode": false,
      "vuex": true,
      "cssPreprocessor": "stylus"
}
EOF

---- ~/.vuerc
{
  "useTaobaoRegistry": false,
  "presets": {
    "my-vue-preset": {
      "useConfigFiles": true,
      "plugins": {
        "@vue/cli-plugin-babel": {},
        "@vue/cli-plugin-eslint": {
          "config": "base",
          "lintOn": [
            "save"
          ]
        },
        "@vue/cli-plugin-unit-jest": {},
        "@vue/cli-plugin-e2e-cypress": {}
      },
      "router": true,
      "routerHistoryMode": false,
      "vuex": true,
      "cssPreprocessor": "stylus"
    }
  }
}

vue create --no-git --preset ./vue-preset.json gut
cd gut
npm run serve
npm run build // for production build

