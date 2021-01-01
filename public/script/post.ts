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
  if (title.innerText.length >= 50) {
    alert("sorry but is too bigger for send that to the server");
  } else if (mineatura.innerText.length >= 100) {
    alert("find another url more small");
  } else if (body.innerText.length >= 100000) {
  }
  console.log("sorry but wtf?");
  console.log("sorry");
  console.log(title.innerText, mineatura.innerText, body.innerText);
  fetch(window.location.pathname, {
    method: "POST",
    body: JSON.stringify({
      title: title.innerText,
      mineatura: mineatura.innerText,
      body: body.innerText,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });
  console.log("change please");
  title.innerText = "";
  mineatura.innerText = "";
  body.innerText = "";
}
