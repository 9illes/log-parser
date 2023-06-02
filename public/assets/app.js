const tabulatorConf = {
  height: "100%",
  progressiveLoad: "load",
  placeholder: "No Data Set",
  columns: [
    { title: "Date", field: "datetime", width: 180, sorter: "date", headerFilter: "input" },
    { title: "Status", field: "status", width: 90, headerFilter: "list", headerFilterParams: { valuesLookup: true, clearable: true }, hozAlign: "center" },
    { title: "Cache", field: "cache", width: 90, headerFilter: "list", headerFilterParams: { valuesLookup: true, clearable: true }, hozAlign: "center" },
    { title: "Method", field: "method", width: 100, headerFilter: "list", headerFilterParams: { valuesLookup: true, clearable: true } },
    { title: "URI", field: "uri", width: 300, headerFilter: "list", headerFilterParams: { valuesLookup: true, clearable: true } },
    { title: "Bot", field: "is_bot", width: 80, formatter: "tickCross", headerFilter: "tickCross", headerFilterParams: { "tristate": true }, headerFilterEmptyCheck: function (value) { return value === null } },
    { title: "Google", field: "is_google_bot", width: 100, formatter: "tickCross", headerFilter: "tickCross", headerFilterParams: { "tristate": true }, headerFilterEmptyCheck: function (value) { return value === null } },
    { title: "UA", field: "user_agent", width: 500, headerFilter: "input" },
    { title: "IP", field: "ip", width: 280, headerFilter: "input" },
  ],
};

const Log = {};

Log.Table = (function (window, undefined) {
  'use strict';
  let _document = window.document,
    table,
    eventSource,
    _init = function (selector, conf) {
      _initTabulator(selector, conf);
      _initEvents();
    },
    _initTabulator = function (selector, conf) {
      table = new Tabulator(selector, conf);
    },
    _initEvents = function () {
      table.on("cellClick", function (e, cell) {
        let clickableColumns = ['ip', 'user_agent', 'uri'];
        if (clickableColumns.includes(cell.getField())) {
          table.setHeaderFilterValue(cell.getField(), cell.getValue());
        }
      });

      _document.getElementById("group-by-ua").addEventListener("click", function () {
        table.setGroupBy('user_agent');
      });

      _document.getElementById("group-by-ip").addEventListener("click", function () {
        table.setGroupBy('ip');
      });

      _document.getElementById("clear-filters").addEventListener("click", function () {
        table.clearHeaderFilter();
        table.clearFilter();
        table.setGroupBy();
      });

      _document.getElementById("button-load-ajax").addEventListener("click", function () {
        _document.getElementById("button-load-ajax").setAttribute('disabled', '');
        let el = _document.getElementById("button-load-ajax");
        el.querySelector('.loading-spinner').classList.remove("is-hidden");
        clearTable();
        loadJournalPage(1);
      });
    },
    clearTable = function () {
      table.clearData();
    },
    loadJournalPage = function (page) {
      fetch('/v1/load?' + new URLSearchParams({
        site: _document.getElementById("site").value,
        type: _document.getElementById("journal_type").value,
        date: _document.getElementById("date").value,
        page: page
      }))
        .then(response => response.json())
        .then(journal => {
          table.addData(journal.entries);
          if (journal.next_page > 0) {
            loadJournalPage(journal.next_page);
          } else {
            // reactivate button
            _document.getElementById("button-load-ajax").querySelector('.loading-spinner').classList.add("is-hidden");
            _document.getElementById("button-load-ajax").removeAttribute('disabled');
          }
        })
        .catch(error => alert("Erreur : " + error));
    },
    _getTable = function () {
      return table;
    };

  return {
    init: _init,
    getTable: _getTable,
    loadJournal: loadJournalPage
  };
})(window);

Log.Table.init("#logs", tabulatorConf);
console.log(Log.Table.getTable());
