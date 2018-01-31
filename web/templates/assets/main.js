/*
 * Utils
 */

Config = {
    apiUrl: 'http://localhost:3000/api/v0.1'
};

Constants = {
    endpointTestsContainerCssClass: 'endpoint-tests',
    loadingContainerCssClass: 'loading-container'
};

Element = {
    findParentElementByClassName: function (element, className) {
        while ((element = element.parentElement) && !element.classList.contains(className));
        return element;
    },

    createLoadingElement: function (args) {
        var elementId = args && args.id ? args.id : Utility.newGuid();
        var div = document.createElement('div');
        var span = document.createElement('span');
        var img = document.createElement('img');
        img.src = 'assets/loading.svg';
        div.id = elementId;
        div.className = Constants.loadingContainerCssClass;
        div.appendChild(span);
        div.appendChild(img);
        return div;
    }
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

Utility = {
    newGuid: function () {
        return Math.random().toString(36).substring(2) + (new Date()).getTime().toString(36);
    },
    // TODO: replace with server-side OS scheduled job
    Timer: {
        initEndpointTestTimer: function (request) {
            this.timerElement = request.timerElement;
            this.offsetMinutes = request.offsetMinutes;
            this.timerElapsedCallback = request.timerElapsedCallback;
            this.timerElapsedCallbackArgs = request.timerElapsedCallbackArgs;
            this.setTimerDeadline();
            this.timerInterval = setInterval(this.runInterval, 1000);
        },
        runInterval: function () {
            self = Utility.Timer;
            var now = new Date().getTime();
            var delta = self.timerDeadline - now;
            var hours = Math.floor((delta % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
            var minutes = Math.floor((delta % (1000 * 60 * 60)) / (1000 * 60));
            var seconds = Math.floor((delta % (1000 * 60)) / 1000);
            self.timerElement.innerHTML = hours + "h " + minutes + "m " + seconds + "s ";

            if (delta < 0) {
                self.setTimerDeadline();
                self.timerElapsedCallback(self.timerElapsedCallbackArgs);
            }
        },
        setTimerDeadline: function (args) {
            var deadlineDate = new Date();
            deadlineDate.setMinutes(deadlineDate.getMinutes() + this.offsetMinutes);
            this.timerDeadline = deadlineDate.getTime();
        }
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
        var testsTab = document.createElement('a');
        testsTab.classList.add('nav-link');
        testsTab.id = 'tests-tab';
        testsTab.href = '#tests';
        testsTab.innerHTML = 'Tests';
        testsTab.setAttribute('data-toggle', 'tab');
        testsTab.setAttribute('role', 'tab');
        var testsTabWrapper = document.createElement('li');
        testsTabWrapper.classList.add('nav-item');
        testsTabWrapper.classList.add('active');
        testsTabWrapper.appendChild(testsTab);
        var incidentsTab = document.createElement('a');
        incidentsTab.classList.add('nav-link');
        incidentsTab.id = 'incidents-tab';
        incidentsTab.href = '#incidents';
        incidentsTab.innerHTML = 'Incidents';
        incidentsTab.setAttribute('data-toggle', 'tab');
        incidentsTab.setAttribute('role', 'tab');
        var incidentsTabWrapper = document.createElement('li');
        incidentsTabWrapper.classList.add('nav-item');
        incidentsTabWrapper.appendChild(incidentsTab);
        var tabBar = document.createElement('ul');
        tabBar.classList.add('nav');
        tabBar.classList.add('nav-tabs');
        tabBar.id = 'endpointTabBar';
        tabBar.role = 'tablist';
        tabBar.appendChild(testsTabWrapper);
        tabBar.appendChild(incidentsTabWrapper);

        var testTabContent = document.createElement('div');
        testTabContent.id = 'tests';
        testTabContent.classList.add('tab-pane');
        testTabContent.classList.add('fade');
        testTabContent.classList.add('active');
        testTabContent.classList.add('in');
        testTabContent.role = 'tabpanel';
        var incidentTabContent = document.createElement('div');
        incidentTabContent.id = 'incidents';
        incidentTabContent.classList.add('tab-pane');
        incidentTabContent.classList.add('fade');
        incidentTabContent.role = 'tabpanel';
        var tabContent = document.createElement('div');
        tabContent.classList.add('tab-content');
        tabContent.id = 'endpointTabContent';
        tabContent.appendChild(testTabContent);
        tabContent.appendChild(incidentTabContent);

        var endpointsViewContainer = document.createElement('div');
        var root = document.createElement('div');
        root.appendChild(endpointsViewContainer);

        EndpointsView.load({
            parent: endpointsViewContainer
        });

        EndpointTestsView.load({
            parent: testTabContent
        });

        IncidentsView.load({
            parent: incidentTabContent
        });

        endpointsViewContainer.appendChild(tabBar);
        endpointsViewContainer.appendChild(tabContent);
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

        this.loadEndpoints({
            parent: endpointsWrapper
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
    }
};

EndpointTestsView = {
    load: function (request) {
        var testDate = document.createElement('th');
        testDate.innerHTML = "Test Date";
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
        var endpointTestsTable = document.createElement('table');
        endpointTestsTable.classList.add("table");
        endpointTestsTable.appendChild(endpointTestsTableHead);
        endpointTestsTable.appendChild(endpointTestsTableBody);

        var endpointTestsMenu = this.createEndpointsTestMenu({
            endpointTestsContainer: endpointTestsTableBody
        });
        var endpointTestsContainer = document.createElement('div');
        endpointTestsContainer.classList.add(Constants.endpointTestsContainerCssClass);
        endpointTestsContainer.appendChild(endpointTestsMenu);
        endpointTestsContainer.appendChild(endpointTestsTable);
        request.parent.appendChild(endpointTestsContainer);

        this.loadEnpointTests({
            parent: endpointTestsTableBody
        });
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
        var self = EndpointTestsView;
        for (var i = 0; i < resp.EndpointTests.length; i++) {
            var row = self.createEndpointTestRow({
                createdAt: resp.EndpointTests[i].CreatedAt,
                name: resp.EndpointTests[i].Name,
                url: resp.EndpointTests[i].URL,
                responseStatus: resp.EndpointTests[i].ResponseStatus,
                timeElapsed: resp.EndpointTests[i].TimeElapsed
            });

            args.parent.appendChild(row);
        }
    },

    runAllEndpointTests: function (request) {
        var self = EndpointTestsView;
        var url = Config.apiUrl + '/endpoints/tests';
        var loading = Element.createLoadingElement();
        request.parent.appendChild(loading);
        Ajax.getJsonRequest({
            url: url,
            callback: self.runAllEndpointTestsCallback,
            callbackArgs: {
                parent: request.parent,
                loadingElement: loading
            }
        });
    },

    runAllEndpointTestsCallback: function (resp, args) {
        var self = EndpointTestsView;
        for (var i = 0; i < resp.length; i++) {
            var row = self.createEndpointTestRow({
                createdAt: resp[i].CreatedAt,
                name: resp[i].Name,
                url: resp[i].URL,
                responseStatus: resp[i].ResponseStatus,
                timeElapsed: resp[i].TimeElapsed
            });

            args.parent.insertBefore(row, args.parent.firstChild);
        }

        if (args.loadingElement) {
            args.parent.removeChild(args.loadingElement);
        }
    },

    createEndpointTestRow: function (request) {
        var testDate = document.createElement('td');
        testDate.innerHTML = request.createdAt;
        var name = document.createElement('td');
        name.innerHTML = request.name;
        var url = document.createElement('td');
        url.innerHTML = request.url;
        var responseStatus = document.createElement('td');
        responseStatus.innerHTML = request.responseStatus;
        var timeElapsed = document.createElement('td');
        timeElapsed.innerHTML = request.timeElapsed;
        var row = document.createElement('tr');
        row.appendChild(testDate);
        row.appendChild(name);
        row.appendChild(url);
        row.appendChild(responseStatus);
        row.appendChild(timeElapsed);
        return row;
    },

    createEndpointsTestMenu: function (args) {
        var endpointTestsRunButton = document.createElement('button');
        endpointTestsRunButton.innerHTML = 'Run Tests';
        endpointTestsRunButton.classList.add('btn');
        endpointTestsRunButton.classList.add('btn-default');
        endpointTestsRunButton.onclick = this.endpointTestsRunButtonEventHandler;
        var endpointTestsButtonWrapper = document.createElement('div');
        endpointTestsButtonWrapper.classList.add('col-md-6');
        endpointTestsButtonWrapper.appendChild(endpointTestsRunButton);
        var endpointTestsTimerWrapper = document.createElement('div');
        endpointTestsTimerWrapper.classList.add('col-md-6');
        endpointTestsTimerWrapper.classList.add('endpoint-tests-timer');
        var nextTestLabel = document.createElement('span');
        nextTestLabel.innerHTML = 'next test';
        nextTestLabel.classList.add('label');
        var nextTestCountdown = document.createElement('span');
        endpointTestsTimerWrapper.appendChild(nextTestLabel);
        endpointTestsTimerWrapper.appendChild(nextTestCountdown);
        Utility.Timer.initEndpointTestTimer({
            // TODO: get from settings
            offsetMinutes: 180,
            timerElement: nextTestCountdown,
            timerElapsedCallback: this.runAllEndpointTests,
            timerElapsedCallbackArgs: {
                parent: args.endpointTestsContainer
            }
        });
        var endpointTestsMenu = document.createElement('div');
        endpointTestsMenu.classList.add('row');
        endpointTestsMenu.classList.add('endpoint-tests-menu');
        endpointTestsMenu.appendChild(endpointTestsButtonWrapper);
        endpointTestsMenu.appendChild(endpointTestsTimerWrapper);
        return endpointTestsMenu;
    },

    endpointTestsRunButtonEventHandler: function (e) {
        var self = EndpointTestsView;
        var endpointTestsContainer = Element.findParentElementByClassName(e.target, Constants.endpointTestsContainerCssClass);
        var endpointTestsBody = endpointTestsContainer.getElementsByTagName('tbody')[0];
        if (!endpointTestsBody) {
            console.log('could not find endpoint tests container');
        } else {
            self.runAllEndpointTests({
                parent: endpointTestsBody
            });
        }
    }
};

IncidentsView = {
    load: function (request) {
        var incidentDate = document.createElement('th');
        incidentDate.innerHTML = "Reported Date";
        var endpointName = document.createElement('th');
        endpointName.innerHTML = "Endpoint";
        var details = document.createElement('th');
        details.innerHTML = "Details";
        var incidentsTableHead = document.createElement('thead');
        incidentsTableHead.appendChild(incidentDate);
        incidentsTableHead.appendChild(endpointName);
        incidentsTableHead.appendChild(details);
        var incidentsTableBody = document.createElement('tbody');
        var incidentsTable = document.createElement('table');
        incidentsTable.classList.add("table");
        incidentsTable.appendChild(incidentsTableHead);
        incidentsTable.appendChild(incidentsTableBody);
        this.loadIncidents({
            parent: incidentsTableBody
        });

        var menu = this.createMenu();
        var root = document.createElement('div');
        root.appendChild(menu);
        root.appendChild(incidentsTable);
        request.parent.appendChild(root);
    },

    createMenu: function () {
        var newButton = document.createElement('button');
        newButton.innerHTML = 'New Incident';
        newButton.classList.add('btn');
        newButton.classList.add('btn-default');
        newButton.onclick = this.newButtonEventHandler;
        var buttonsWrapper = document.createElement('div');
        buttonsWrapper.classList.add('col-md-6');
        buttonsWrapper.appendChild(newButton);
        var otherActionsWrapper = document.createElement('div');
        otherActionsWrapper.classList.add('col-md-6');
        var menu = document.createElement('div');
        menu.classList.add('row');
        menu.classList.add('incidents-menu');
        menu.appendChild(buttonsWrapper);
        menu.appendChild(otherActionsWrapper);
        return menu;
    },

    loadIncidents: function (request) {
        var url = Config.apiUrl + '/endpoints/incidents';
        Ajax.getJsonRequest({
            url: url,
            callback: this.loadIncidentsCallback,
            callbackArgs: {
                parent: request.parent
            }
        });
    },

    loadIncidentsCallback: function (res, args) {
        var self = IncidentsView;
        for (var i = 0; i < res.Incidents.length; i++) {
            var row = self.createIncidentRow({
                createdAt: res.Incidents[i].IncidentCreatedAt,
                endpointName: res.Incidents[i].EndpointName,
                urgency: res.Incidents[i].IncidentUrgency,
                impact: res.Incidents[i].IncidentImpact,
                details: res.Incidents[i].IncidentDetails
            });

            args.parent.appendChild(row);
        }
    },

    createIncidentRow: function (args) {
        var incidentDate = document.createElement('td');
        incidentDate.innerHTML = args.createdAt;
        var endpointName = document.createElement('td');
        endpointName.innerHTML = args.endpointName;
        var details = document.createElement('td');
        details.innerHTML = args.details;
        var row = document.createElement('tr');
        row.appendChild(incidentDate);
        row.appendChild(endpointName);
        row.appendChild(details);
        return row;
    },

    newButtonEventHandler: function (e) {}
};

/*
 * Main
 */

RootView.load({
    parent: document.getElementById('root')
});