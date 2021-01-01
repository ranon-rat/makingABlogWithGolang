function sendDocument(): void {
  let [title, mineatura, body]: HTMLElement[] = [
    document.getElementById("titulo"),
    document.getElementById("mineatura"),
    document.getElementById("body"),
  ];
  if (!title.innerText || !mineatura.innerText || !body.innerText) {
    alert("you need to define everything");
    return;
  }
  if (
    title.innerText.length >= 50 ||
    mineatura.innerText.length >= 100 ||
    body.innerText.length >= 100000
  ) {
    alert("sorry but is too bigger for send that to the server");
  }

  fetch(window.location.pathname, {
    method: "POST",
    body: JSON.stringify({
      title: title.innerText,
      mineatura: mineatura.innerText,
      bodyOfDocument: body.innerText,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });
  console.log(title.innerText, mineatura.innerText, body.innerText);

  title.innerText = "";
  mineatura.innerText = "";
  body.innerText = "";
}
