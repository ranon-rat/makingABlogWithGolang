// @ts-nocheck
interface Formulario {
  title: any;
  mineatura: any;
  bodyOfDocument: any;
}

async function sendDocument() {
  let f: Formulario = {
    title: document.getElementById("titulo")?.innerText,
    mineatura: document.getElementById("mineatura")?.innerText,
    bodyOfDocument: document.getElementById("bodyOfMessage")?.value,
  };

  if (!f.title || !f.mineatura || !f.bodyOfDocument) {
    alert("you need to define everything");
    return;
  }
  if (f.bodyOfDocument.length >= 100000) {
    alert("sorry but is too bigger for send that to the server");
    return;
  }

  fetch(window.location.pathname, {
    method: "POST",
    body: JSON.stringify(f),
    headers: {
      "Content-Type": "application/json",
    },
  });
  document.getElementById("titulo")?.innerText = "";
  document.getElementById("mineatura")?.innerText = "";
  document.getElementById("bodyOfMessage")?.value = "";

  console.log(f);
}
