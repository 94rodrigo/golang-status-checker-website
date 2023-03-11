const tooltipTriggerList = () => document.querySelectorAll('[data-bs-toggle="tooltip"]');
const tooltipList = () => [...tooltipTriggerList()].map(tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl));
const toastStatusUpdated = () => document.getElementById('toast-status-update');
const toastMessage = () => document.querySelector('#toast-message');

function tooltipInit() {
    tooltipTriggerList();
    tooltipList();
}

function getUpdateStatus(idWebsite) {
    fetch('/' + idWebsite)
        .then(response => response.json())
        .then(data => {
            console.log(data);
            let columnSelector = document.querySelector('#icn-circle-' + idWebsite);
            if (data.reqStatus <= 399)
                columnSelector.setAttribute('src', 'https://img.icons8.com/emoji/48/null/green-circle-emoji.png');
            else
                columnSelector.setAttribute('src', 'https://img.icons8.com/emoji/48/null/red-circle-emoji.png');

            columnSelector.setAttribute('data-bs-title', data.reqStatusDesc);
            tooltipInit();
            toastMessage().textContent = 'Status successfully updated';
            new bootstrap.Toast(toastStatusUpdated()).show();
        })
        .catch(() => {
            toastMessage().textContent = 'ERRROR: please, try again';
            new bootstrap.Toast(toastStatusUpdated()).show();
        });
}

function deleteWebsite(idWebsite) {
    fetch('/remove/' + idWebsite, { method: 'DELETE', redirect: 'error' })
        .then(response => response.json())
        .then(data => console.log(data))
        .then(() => setTimeout(refreshPage(), 1000));
}
    
function refreshPage() {
    history.go(0);
    window.location.reload();
    window.location.href = window.location.href;
}

function verifyErrorInJsonResponse() {
    let jsonResponseString = document.querySelector('#json-response').textContent;
    let input = document.querySelector('#ipt-website-field');
    let validationServerWebsiteFeedback = document.querySelector('#validationServerWebsiteFeedback');
    if (jsonResponseString != '') {
        let parsedJsonResponse = JSON.parse(jsonResponseString)
        if (parsedJsonResponse.error && parsedJsonResponse.error != '') {
            input.classList.add('is-invalid');
            validationServerWebsiteFeedback.textContent = parsedJsonResponse.error;
            validationServerWebsiteFeedback.removeAttribute('hidden');
        } else {
            input.classList.remove('is-invalid');
            validationServerWebsiteFeedback.setAttribute('hidden', 'hidden');
        }
    } else {
        input.classList.remove('is-invalid');
        validationServerWebsiteFeedback.setAttribute('hidden', 'hidden');
    }
}

(function () {
    tooltipInit();
    verifyErrorInJsonResponse();
})();