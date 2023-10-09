window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
      {
        url: "lancr/v1/hero_service.swagger.json",
        name: "Hero Service",
      },
      {
        url: "lancr/v1/apprentice_service.swagger.json",
        name: "Apprentice Service",
      },
      {
        url: "lancr/v1/quest_service.swagger.json",
        name: "Quest Service",
      },
    ],
    dom_id: "#swagger-ui",
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: "StandaloneLayout",
  });

  //</editor-fold>
};
