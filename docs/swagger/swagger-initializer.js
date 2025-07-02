window.onload = () => {
  window.ui = SwaggerUIBundle({
    urls: [
      { url: "auth.json", name: "auth Service" },
    ],
    dom_id: "#swagger-ui",
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    layout: "StandaloneLayout"
  });
};
