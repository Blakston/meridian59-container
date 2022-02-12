// createConfig will handle the low-level communication configuration.
const createConfig = (method, params) => {
    return {
        method: method,
        headers: {
            "accept": "application/json",
            "content-type": "application/json"
        },
        body: JSON.stringify(params)
    };
};

// Component introduces a state which can be accessed via a get/set.
// Additionally it provides event handling by receiving events via on and emitting events via emit.
class Component {
    // At first we construct an HTML element and set the state to an empty object.
    constructor() {
        this.state = {};
    }
    // emit dispatches a specific event with corresponding data.
    emit(event, data) {
        // We use window for dispatching events globally.
        // Thus, we don't need "bubbles" to propagate events up through the DOM.
        window.dispatchEvent(new CustomEvent(event, {
            detail: {
                output: data
            }
        }))
    }
    // getState reads a state value by a given key.
    getState(key) {
        return this.state[key];
    }
    // on adds a listener for a specific event.
    on(event, fn) {
        window.addEventListener(event, (e) => {
            // Only the object data (detail) is necessary for this kind of event.
            fn(e.detail.output);
        });
    }
    // setState writes a state key, value pair.
    setState(key, val) {
        this.state[key] = val;
    }
}

// Model handles the business logic by using generated functions.
class Model extends Component {
    // Constructor ...
    constructor() {
        super();
        // Add event listeners
        this.on("Show", async (params) => {
            let config = createConfig("POST", params);
            fetch("http://127.0.0.1/api/accounts/show", config)
                .then((response) => response.json())
                .then((data) => {
                    this.emit("Show done", data);
                })
                .catch((err) => {
                    this.emit("Show error", err);
                })
        });
    }
}

// ViewModel handles the state and provides a simple API.
class ViewModel extends Component {
    // Constructor ...
    constructor() {
        super();
        // Add event listeners
        this.on("Show done", (data) => {
            this.setState("Show result", data);
        });
    }
    // Create API calls	
    Show(data) {
        this.emit("Show", data);
    }
}

// View ...
class View extends Component {
    // Constructor ...
    constructor(viewModel) {
        super();
        this.viewModel = viewModel;
        // Add event listeners
        this.on("Show done", (data) => {
            this.render();
        });
        this.on("Show error", (err) => {
            this.render();
        });
        // Initial rendering
        this.render();
    }
    // render ...
    render() {
        // Read the current state
        let show = this.viewModel.getState("Show result");
        // Set the default
        if (typeof show === "undefined") {
            show = []; // default here
        }
        // Set contents
        let showHtml = document.querySelector("#online");
        showHtml.innerHTML = "";
        show.forEach(toonName => {
            let p = document.createElement("p");
            p.innerText = toonName; 
            showHtml.appendChild(p);
        });
        document.querySelector("#online").innerHTML = show;
        // Add DOM event listeners
        let self = this;
        document.querySelector("#create").addEventListener("click", (evt) => {
            evt.preventDefault();
            document.querySelector("#user").value = "";
            document.querySelector("#pass").value = "";
            document.querySelector("#pass2").value = "";
        });
        // Add interval for who list
        if (!reloadInterval) {
            reloadInterval = setInterval(() => {
                self.viewModel.Show();
            }, 10000)
        }
    }
}

let reloadInterval;

const model = new Model();
const viewModel = new ViewModel(); // model is loosely coupled via events
const view = new View(viewModel);