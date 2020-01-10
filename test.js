var CryptoJS = require('crypto-js');
var request = require('request');

var reqJson = {
  Customer: {
    //Id: null,
    Name: 'FirstCustomer Inc',
    RegNo: '1122334455',
    NotTDCustomer: false,
    VatRegNo: '11223344',
    CurrencyCode: 'EUR',
    PaymentDeadLine: 7,
    OverDueCharge: 0,
    RefNoBase: 1,
    Address: 'Merimiehenkatu 31',
    CountryCode: 'FI',
    County: 'Finland',
    City: 'Helsinki',
    PostalCode: '',
    PhoneNo: '6548765',
    PhoneNo2: '',
    HomePage: '',
    Email: 'customermail@gmail.com',
  },
  DocDate: '20170113131239',
  DueDate: '20170125131239',
  InvoiceNo: '123',
  RefNo: '222',
  DepartmentCode: '',
  ProjectCode: '',
  InvoiceRow: [
    {
      Item:
      {
        Code: 1234567,
        Description: 'Bag of goldflakes',
        Type: 3,
        UOMName: 'kg',
      },
      Quantity: '2',
      Price: '1000',
      DiscountPct: 0,
      DiscountAmount: 0,
      TaxId: '665f01a4-357a-4a6b-a565-2f17e6e1da13',
      LocationCode: 1,
    },
  ],
  TotalAmount: 2000,
  RoundingAmount: 0,
  TaxAmount: [{ TaxId: '665f01a4-357a-4a6b-a565-2f17e6e1da13', Amount: 400 }],
  HComment: '',
  FComment: '',
};

var ApiId = process.env.API_ID;
var ApiKey = process.env.API_KEY;

function pad2(n) {
  return n > 9 ? '' + n : '0' + n;
}

function getTimestamp() {
  var d = new Date();
  var yyyy = d.getFullYear();
  var MM = pad2(d.getMonth() + 1);
  var dd = pad2(d.getDate());
  var HH = pad2(d.getHours());
  var mm = pad2(d.getMinutes());
  var ss = pad2(d.getSeconds());
  return yyyy + MM + dd + HH + mm + ss;
}

var timestamp = getTimestamp();
var dataString = ApiId + timestamp + JSON.stringify(reqJson);
var hash = CryptoJS.HmacSHA256(dataString, ApiKey);
var signature = CryptoJS.enc.Base64.stringify(hash);
console.log(dataString);
console.log(hash);
console.log(signature);
var url = 'https://aktiva.merit.ee/api/v1/sendinvoice' +
'?ApiId=' + ApiId + '&timestamp=' + timestamp + '&signature=' +
  signature;
request({
  url: url,
  method: 'POST',
  json: true,
  headers: { 'content-type': 'application/json', },
  body: reqJson,
},
  function (request, response) {
    console.log('Status code: ', response.statusCode, ' -- ',
      response.statusMessage);
    console.log('Headers: ', response.headers['content-type']);
    console.log('Body: ', response.body);
  }
);
