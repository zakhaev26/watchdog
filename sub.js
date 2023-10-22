console.log("connected!")
const source = new EventSource('http://localhost:8080/stream');
const container = document.getElementById("sse");
source.onmessage = (e) =>{
    console.log("Received !");
    const data = JSON.parse(e.data)
    console.log(data);
    const ele = document.createElement("h1");
    ele.innerText = `Event Num : ${data}`;
    container.appendChild(ele);
}

source.onerror = (e) => {
    if (e.eventPhase === EventSource.CLOSED) {
        const end = document.createElement("h1");
        end.innerText = `Publisher/Die`;
        container.appendChild(end);
    }
};


/**
 *
{ isTrusted: true
    bubbles: false
    cancelBubble: false
    cancelable: false
    composed: false
    currentTarget: EventSource {url: 'http://localhost:8080/stream', withCredentials: false, readyState: 1, onopen: null, onmessage: ƒ, …}
    data: "19"
    defaultPrevented: false
    eventPhase: 0
    lastEventId: ""
    origin: "http://localhost:8080"
    ports: []
    returnValue: true
    source: null
    srcElement: EventSource {url: 'http://localhost:8080/stream', withCredentials: false, readyState: 1, onopen: null, onmessage: ƒ, …}
    target: 
    EventSource {url: 'http://localhost:8080/stream', withCredentials: false, readyState: 1, onopen: null, onmessage: ƒ, …}
    timeStamp: 19100.59999999404
    type: "message"
    userActivation: null
    [[Prototype]]: MessageEvent
}
 */