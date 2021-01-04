var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
function NewPublications() {
    return __awaiter(this, void 0, void 0, function () {
        var urlApi, publication, d, _i, _a, i, element, pagePublications, i, Element_1;
        return __generator(this, function (_b) {
            switch (_b.label) {
                case 0:
                    urlApi = "api" +
                        "/" +
                        window.location.pathname.split("/")[window.location.pathname.length - 1];
                    publication = {
                        Publications: [
                            {
                                id: 0,
                                title: "",
                                mineatura: "",
                                body: ""
                            },
                        ],
                        Size: 0,
                        Cantidad: 0
                    };
                    // this is only for get the api
                    return [4 /*yield*/, fetch(urlApi)
                            .then(function (r) { return r.text(); })
                            .then(function (data) {
                            console.log(data);
                            publication = data;
                            publication = JSON.parse(publication);
                        })];
                case 1:
                    // this is only for get the api
                    _b.sent();
                    d = document.getElementById("publications");
                    for (_i = 0, _a = publication.Publications; _i < _a.length; _i++) {
                        i = _a[_i];
                        element = "\n    <p>\n    <a  class=\"publications\" href=\"/publication/" + i.id + "\">\n      <div >\n        <div class=\"publicationContent\">\n        <img src=\"" + i.mineatura + "\"style=\"height: 8em;\" > \n              \n        <h4 >\n            " + i.title + "\n          </h4>\n         \n          \n          </div>\n        </div>\n      </div>\n    </a>\n    </p>\n      \n    ";
                        d.innerHTML += element;
                    }
                    pagePublications = document.getElementById("pagePublications");
                    for (i = publication.Size / publication.Cantidad; i > 0; i--) {
                        Element_1 = "\n \n    <div class=\"buttonElementID\" style=\"background-color: rgb(255, 255, 255);\n     width:1em;\n    height: 1em;\">\n      <a class=\"buttonElementID\" style=\"background-color: rgb(255, 255, 255);\" href=\"/" + Math.round(publication.Size / publication.Cantidad - i) + "\" style=\" >\n        <div class=\"buttonElementID\" style=\"  \n        \n        text-decoration: none;\n        color:rgb(0, 0, 0);\n        display:inline-block;\n        width:10px;\n        height: 10px;\n        \" >\n          <h4 > " + (Math.round(publication.Size / publication.Cantidad - i) + 1) + " </h4>\n        </div>\n      </a>\n    </div >\n    ";
                        pagePublications.innerHTML += Element_1;
                        console.log(Math.round(publication.Size / 10 - i) + 1);
                    }
                    console.log(publication);
                    return [2 /*return*/];
            }
        });
    });
}
NewPublications();
