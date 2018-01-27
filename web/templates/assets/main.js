/*
 * Utils
 */

Config = {
    apiUrl: 'http://localhost:3000/api/v0.1'
};

Ajax = {
    getJsonRequest: function (args) {
        this.getRequest({
            url: args.url,
            contentType: 'application/json',
            callback: args.callback,
            callbackArgs: args.callbackArgs
        });
    },

    postJsonRequest: function (args) {
        this.postRequest({
            url: args.url,
            contentType: 'application/json',
            content: args.content ? JSON.stringify(args.content) : '',
            callback: args.callback,
            callbackArgs: args.callbackArgs
        });
    },

    getRequest: function (args) {
        this.request({
            url: args.url,
            method: 'GET',
            contentType: args.contentType,
            callback: args.callback,
            callbackArgs: args.callbackArgs
        });
    },

    postRequest: function (args) {
        this.request({
            url: args.url,
            method: 'POST',
            contentType: args.contentType,
            content: args.content,
            callback: args.callback,
            callbackArgs: args.callbackArgs
        });
    },

    request: function (args) {
        var httpRequest = new XMLHttpRequest();
        if (!('withCredentials' in httpRequest)) httpRequest = new XDomainRequest(); // fix IE8/9

        if (!httpRequest) {
            this.logError('Could not create ajax http request');
            return false;
        }

        httpRequest.onreadystatechange = function () {
            Ajax.response(httpRequest, args.callback, args.callbackArgs);
        };
        httpRequest.open(args.method, args.url);
        httpRequest.withCredentials = true;

        if (args.contentType) {
            httpRequest.setRequestHeader("Content-Type", args.contentType);
        }

        if (args.authorization) {
            httpRequest.setRequestHeader("Authorization", args.authorization);
        }

        httpRequest.send(args.content);
    },

    response: function (httpRequest, callbackFunction, callbackArgs) {
        if (httpRequest.readyState === XMLHttpRequest.DONE && callbackFunction) {
            if (httpRequest.status === 200 && httpRequest.responseText) {
                // TODO: check if content-type == application/json before converting
                callbackFunction(JSON.parse(httpRequest.responseText), callbackArgs);
            } else {
                Ajax.logError('{ "status": ' + httpRequest.status + ', "responseText": "' + httpRequest.responseText + '" }');
            }
        }
    },

    logError: function (message) {
        console.log(message);
    }
};

/*
 * Views
 */

RootView = {
    load: function (request) {
        NavBarView.load(request);
        ContentView.load(request);
    }
};

NavBarView = {
    load: function (request) {
        var root = this.createRootElement();
        request.parent.appendChild(root);
    },

    createRootElement: function () {
        var toggleNavElement = this.createToggleNavElement();
        var navBrandElement = this.createNavBrandElement();
        var beta = document.createElement('sup');
        beta.className = 'navbar-brand beta';
        beta.innerHTML = 'BETA';
        var navHeader = document.createElement('div');
        navHeader.className = 'navbar-header';
        navHeader.appendChild(toggleNavElement);
        navHeader.appendChild(navBrandElement);
        // navHeader.appendChild(beta);
        var nav = this.createNavBarElement();
        var navWrapper = document.createElement('div');
        navWrapper.className = 'container-fluid';
        navWrapper.appendChild(navHeader);
        navWrapper.appendChild(nav);
        var nav = document.createElement('div');
        nav.className = 'navbar navbar-inverse navbar-fixed-top';
        nav.appendChild(navWrapper);
        return nav;
    },

    createToggleNavElement: function () {
        var toggleNav = document.createElement('span');
        toggleNav.className = 'sr-only';
        toggleNav.innerHTML = 'Toggle navigation';
        var iconBar1 = document.createElement('span');
        iconBar1.className = 'icon-bar';
        var iconBar2 = document.createElement('span');
        iconBar2.className = 'icon-bar';
        var iconBar3 = document.createElement('span');
        iconBar3.className = 'icon-bar';
        var toggleNavButton = document.createElement('button');
        toggleNavButton.className = 'navbar-toggle collapsed';
        toggleNavButton.type = 'button';
        toggleNavButton.setAttribute('data-toggle', 'collapse');
        toggleNavButton.setAttribute('data-target', '#navbar');
        toggleNavButton.setAttribute('aria-expanded', 'false');
        toggleNavButton.setAttribute('aria-controls', 'navbar');
        toggleNavButton.appendChild(toggleNav);
        toggleNavButton.appendChild(iconBar1);
        toggleNavButton.appendChild(iconBar2);
        toggleNavButton.appendChild(iconBar3);
        return toggleNavButton;
    },

    createNavBrandElement: function () {
        var navBrandName = document.createElement('p');
        var name = document.createElement('span');
        // TODO: pull name text from request <- api
        name.innerHTML = 'ETS System Dashboard';
        name.style.color = '#dddddd';
        var navBrand = document.createElement('span');
        navBrand.className = 'navbar-brand';
        navBrand.appendChild(name);
        return navBrand;
    },

    createNavBarElement: function () {
        var navbar = document.createElement('div');
        navbar.className = 'collapse navbar-collapse';
        return navbar;
    }
};

ContentView = {
    load: function (request) {
        var root = this.createRootElement();
        request.parent.appendChild(root);
    },

    createRootElement: function () {
        var endpointsViewContainer = document.createElement('div');
        var root = document.createElement('div');
        root.appendChild(endpointsViewContainer);
        EndpointsView.load({
            parent: endpointsViewContainer
        });
        return root;
    }
};

EndpointsView = {
    load: function (request) {
        var endpointsWrapper = document.createElement('div');
        endpointsWrapper.classList.add('row');
        var endpointsContainer = document.createElement('div');
        endpointsContainer.appendChild(endpointsWrapper);
        request.parent.appendChild(endpointsContainer);
        var testDate = document.createElement('th');
        testDate.innerHTML = "Test Date &amp; Time";
        var name = document.createElement('th');
        name.innerHTML = "Name";
        var url = document.createElement('th');
        url.innerHTML = "URL";
        var responseStatus = document.createElement('th');
        responseStatus.innerHTML = "Response Status";
        var timeElapsed = document.createElement('th');
        timeElapsed.innerHTML = "Time Elapsed";
        var endpointTestsTableHead = document.createElement('thead');
        endpointTestsTableHead.appendChild(testDate);
        endpointTestsTableHead.appendChild(name);
        endpointTestsTableHead.appendChild(url);
        endpointTestsTableHead.appendChild(responseStatus);
        endpointTestsTableHead.appendChild(timeElapsed);
        var endpointTestsTableBody = document.createElement('tbody');
        var endpointTestsContainer = document.createElement('table');
        endpointTestsContainer.classList.add("table");
        endpointTestsContainer.classList.add("endpoint-tests");
        endpointTestsContainer.appendChild(endpointTestsTableHead);
        endpointTestsContainer.appendChild(endpointTestsTableBody);
        request.parent.appendChild(endpointTestsContainer);

        this.loadEndpoints({
            parent: endpointsWrapper
        });
        this.loadEnpointTests({
            parent: endpointTestsTableBody
        });
    },

    loadEndpoints: function (request) {
        var url = Config.apiUrl + '/endpoints';
        Ajax.getJsonRequest({
            url: url,
            callback: this.loadEndpointsCallback,
            callbackArgs: {
                parent: request.parent
            }
        });
    },

    loadEndpointsCallback: function (resp, args) {
        for (var i = 0; i < resp.length; i++) {
            var name = document.createElement('div');
            name.classList.add('name');
            name.innerHTML = resp[i].Name;
            var url = document.createElement('div');
            url.classList.add('url');
            url.innerHTML = resp[i].URL;
            var item = document.createElement('div');
            item.classList.add('col-sm-2');
            item.classList.add('col-md-4');
            item.classList.add('endpoint');
            item.appendChild(name);
            item.appendChild(url);
            args.parent.appendChild(item);
        }
    },

    loadEnpointTests: function (request) {
        var url = Config.apiUrl + '/endpoint-tests';
        Ajax.getJsonRequest({
            url: url,
            callback: this.loadEndpointTestsCallback,
            callbackArgs: {
                parent: request.parent
            }
        });
    },

    loadEndpointTestsCallback: function (resp, args) {
        for (var i = 0; i < resp.EndpointTests.length; i++) {
            var testDate = document.createElement('td');
            testDate.innerHTML = resp.EndpointTests[i].CreatedAt;
            var name = document.createElement('td');
            name.innerHTML = resp.EndpointTests[i].Name;
            var url = document.createElement('td');
            url.innerHTML = resp.EndpointTests[i].URL;
            var responseStatus = document.createElement('td');
            responseStatus.innerHTML = resp.EndpointTests[i].ResponseStatus;
            var timeElapsed = document.createElement('td');
            timeElapsed.innerHTML = resp.EndpointTests[i].TimeElapsed;
            var row = document.createElement('tr');
            row.appendChild(testDate);
            row.appendChild(name);
            row.appendChild(url);
            row.appendChild(responseStatus);
            row.appendChild(timeElapsed);
            args.parent.appendChild(row);
        }
    }
};

/*
 * Main
 */

RootView.load({
    parent: document.getElementById('root')
});