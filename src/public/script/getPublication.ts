interface Pub {
  Cantidad: number;
  Publications: {
    id: number;
    title: string;
    mineatura: string;
    bodyOfDocument: string;
  }[];
  Size: number;
}

async function NewPublications() {
  let urlApi: string =
    "api" + "/" + window.location.pathname.replace(/\//gi, "");
  let publication: Pub | string = {
    Publications: [
      {
        id: 0,
        title: "",
        mineatura: "",
        bodyOfDocument: "",
      },
    ],
    Size: 0,
    Cantidad: 0,
  };

  // this is only for get the api
  await fetch(urlApi)
    .then((r) => r.text())
    .then((data) => {
      publication = data;
      publication = JSON.parse(publication);
    });

  let d: any = document.getElementById("publications");

  // add the elements
  for (let i of publication.Publications) {
    // then add elements into the dom
    let element = `
    <p>
      <a  class="publications" href="/publication/${i.id}">
        <div >
          <div class="publicationContent">
          <img src="${i.mineatura}" >    
          <div class="aboutPubication"> 
            <h4 >
              ${i.title}
            </h4>
            <p >
            ${i.bodyOfDocument.slice(0, 30).replace(/#|'|`|"|\||-|@|=/gi, "")}
            </p>
          </div>
            
            </div>
          </div>
        </div>
      </a>
    </p>
      
    `;

    d.innerHTML += element;
  }
  let pagePublications: any = document.getElementById("pagePublications");
  for (
    let i: number = 1;
    i <= publication.Size / publication.Cantidad + 1;
    i++
  ) {
    let Element: string = `
 
    <div class="buttonElementID" style="background-color: rgb(255, 255, 255);
     width:1em;
    height: 1em;">
      <a class="buttonElementID" style="background-color: rgb(255, 255, 255);" href="/${i}" style=" >
        <div class="buttonElementID" style="  
        
        text-decoration: none;
        color:rgb(0, 0, 0);
        display:inline-block;
        width:10px;
        height: 10px;
        " >
        <h4 > ${i} </h4>
        </div>
      </a>
    </div >
    `;

    pagePublications.innerHTML += Element;
    console.log(i); // this supose to be you
  }
}
NewPublications();
