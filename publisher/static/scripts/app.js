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
        this.on("Create", async (params) => {
            let config = createConfig("POST", params);
            fetch("/api/accounts/create", config)
                .then((response) => response.json())
                .then((data) => {
                    this.emit("Create done", data);
                })
                .catch((err) => {
                    this.emit("Create error", err);
                })
        });
        this.on("Log", async (params) => {
            let config = createConfig("POST", params);
            fetch("/api/god/log", config)
                .then((response) => response.json())
                .then((data) => {
                    this.emit("Log done", data);
                })
                .catch((err) => {
                    this.emit("Log error", err);
                })
        });
        this.on("Online", async (params) => {
            let config = createConfig("POST", params);
            fetch("/api/accounts/online", config)
                .then((response) => response.json())
                .then((data) => {
                    this.emit("Online done", data);
                })
                .catch((err) => {
                    this.emit("Online error", err);
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
        this.on("Create done", (data) => {
            this.setState("Create result", data);
        });
        this.on("Log done", (data) => {
            this.setState("Log result", data);
        });
        this.on("Online done", (data) => {
            this.setState("Online result", data);
        }); 
    }
    // Create API calls	
    Create(data) {
        this.emit("Create", data);
    }
    Log(data) {
        this.emit("Log", data);
    }
    Show(data) {
        this.emit("Online", data);
    }
}

// View ...
class View extends Component {
    // Constructor ...
    constructor(viewModel) {
        super();
        this.viewModel = viewModel;
        // Add event listeners
        this.on("Create done", (data) => {
            this.render();
        });
        this.on("Create error", (err) => {
            this.render();
        });
        this.on("Log done", (data) => {
            this.render();
        });
        this.on("Log error", (err) => {
            this.render();
        });
        this.on("Online done", (data) => {
            this.render();
        });
        this.on("Online error", (err) => {
            this.render();
        });
        // Initial rendering
        this.render();
    }
    // render ...
    render() {
        // Read the current state
        let create = this.viewModel.getState("Create result");
        let log = this.viewModel.getState("Log result");
        let online = this.viewModel.getState("Online result");
        // Set the default
        if (typeof create === "undefined") {
            create = {"error": ""}; // default here
        }
        if (typeof log === "undefined") {
            log = []; // default here
        }
        if (typeof online === "undefined") {
            online = []; // default here
        }
        // Set contents
        let showStatus = document.querySelector("#account_status");
        showStatus.innerHTML = create.error;
        let showOnline = document.querySelector("#online");
        showOnline.innerHTML = "";
        for (let i = 0; i < online.length; i++) {
            let p = document.createElement("p");
            p.innerText = online[i]; 
            showOnline.appendChild(p);
        }
        let showLog = document.querySelector("#logfile");
        showLog.innerHTML = "";
        for (let i = 0; i < log.length; i++) {
            let p = document.createElement("p");
            p.innerText = log[i]; 
            showLog.appendChild(p);
        }
        // Add DOM event listeners
        let self = this;
        document.querySelector("#create").addEventListener("click", (evt) => {
            evt.preventDefault();
            self.viewModel.Create({
                "user": document.querySelector("#user").value,
                "pass": document.querySelector("#pass").value,
                "pass2": document.querySelector("#pass2").value,
                "email": document.querySelector("#email").value
            });
            document.querySelector("#user").value = "";
            document.querySelector("#pass").value = "";
            document.querySelector("#pass2").value = "";
            document.querySelector("#email").value = "";
        });
        // Add interval for who list
        if (!reloadInterval) {
            reloadInterval = setInterval(() => {
                self.viewModel.Log();
                self.viewModel.Show();
            }, 10000)
        }
    }
}

let reloadInterval;

const model = new Model();
const viewModel = new ViewModel(); // model is loosely coupled via events
const view = new View(viewModel);