function showInfo(message) {
  Toastify({
    text: message,
    duration: 3000,
    newWindow: true,
    gravity: "top", 
    position: "center", 
    style: {
      background: "#6b21a8",
    },
  }).showToast();
}

function showError(message) {
  Toastify({
    text: message,
    duration: 3000,
    newWindow: true,
    gravity: "top", 
    position: "center", 
    style: {
      background: "#be185d",
    },
  }).showToast();
}
