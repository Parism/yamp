var constdata;
var state = 0;

$(document).ready(function() {
    $("#dynform").submit(function(e){
        e.preventDefault(e);
        $("#result").html("Παρακαλώ περιμένετε");
        var date = $("#datehidden").val();
        var label = $("#label").val();
        $.ajax({
            url: "/getdyn",
            data: {
                date:date,
                label:label,
            },
            accepts: {
                text: "application/json; charset=utf-8"
            },
            error: function() {
                $("#result").html("Σφάλμα κατά την ανάκτηση");
              },
            success: function(data){
                constdata = data
                parseDataMin(constdata)
            },
            timeout: 10000 // sets timeout to 10 seconds
        });
    });
});

function parseDataFull(data){
    $("#result").html("")
    $("#result").append("<button id=\"datatoggle\">Συνοπτική</button>")
    $("#result").append("<h3>Παρόντες</h3>")
    if(constdata.proswpiko == null){
        $("#result").append("<h3>Κανείς παρών</h3>")    
    }else{
        unique_ranks = getUniqueValuesOfKey(constdata.proswpiko,"rank")
        for(var rank in unique_ranks){
            var temp = data.proswpiko.filter(obj => {
                return obj.rank === unique_ranks[rank]
            })
            $("#result").append("<h4>"+unique_ranks[rank]+"</h4>")
            ulelement = document.createElement('ul')
            for (var i in temp){
                lielement = document.createElement('li')
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+temp[i]);
                a.innerHTML = temp[i].surname+" "+temp[i].name;
                span1.appendChild(a)
                lielement.appendChild(span1)
                ulelement.appendChild(lielement)
            }
            $("#result").append(ulelement)
        }
    }
    if (constdata.metaboles == null){
        $("#result").append("<h3>Καμία μεταβολή</h3>")
    }else{
        $("#result").append("<h3>Μεταβολές</h3>")
        unique_categories = getUniqueValuesOfKey(constdata.metaboles,"Category")
        for (var category in unique_categories){
            var temp = data.metaboles.filter(obj => {
                return obj.Category === unique_categories[category]
            })
            $("#result").append("<h4>"+unique_categories[category]+"</h4>")
            ulelement = document.createElement('ul')
            for (var i in temp){
                lielement = document.createElement('li')
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+temp[i].PersonID);
                a.innerHTML = temp[i].Surname+" "+temp[i].Name+" ";
                span1.appendChild(a)
                span2 = document.createElement('span')
                span2.innerHTML = temp[i].Repr
                lielement.appendChild(span1)
                lielement.appendChild(span2)
                ulelement.appendChild(lielement)
            }
            $("#result").append(ulelement)
        } 
      }
    if (constdata.ypiresies == null){
        $("#result").append("<h3>Καμία υπηρεσία</h3>")
    }else{
        $("#result").append("<h3>Υπηρεσίες</h3>")
        ulelement = document.createElement('ul')
        for(var ypiresia in constdata.ypiresies){
            lielement = document.createElement('li')
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+constdata.ypiresies[ypiresia].idperson);
                a.innerHTML = constdata.ypiresies[ypiresia].surname+" "+constdata.ypiresies[ypiresia].name;
                span1.appendChild(a)
                span2 = document.createElement('span')
                span2.innerHTML = " "+constdata.ypiresies[ypiresia].rank+" "+constdata.ypiresies[ypiresia].perigrafi
                lielement.appendChild(span1)
                lielement.appendChild(span2)
                ulelement.appendChild(lielement)
        }
        $("#result").append(ulelement)
    }
    if (constdata.aitiseis == null){
        $("#result").append("<h3>Καμία αίτηση</h3>")
    }else{
        $("#result").append("<h3>Αιτήσεις</h3>")
        ulelement = document.createElement('ul')
        for(var aitisi in constdata.aitiseis){
            lielement = document.createElement('li')
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+constdata.aitiseis[aitisi].idperson);
                a.innerHTML = constdata.aitiseis[aitisi].surname+" "+constdata.aitiseis[aitisi].name;
                span1.appendChild(a)
                span2 = document.createElement('span')
                span2.innerHTML = " "+constdata.aitiseis[aitisi].perigrafi
                lielement.appendChild(span1)
                lielement.appendChild(span2)
                ulelement.appendChild(lielement)
        }
        $("#result").append(ulelement)
    }
    if(constdata.anafores == null){
        $("#result").append("<h3>Καμία αναφορά</h3>")
    }else{
        for(var anafora in constdata.anafores){
            lielement = document.createElement('li')
                span1 = document.createElement('span')
                a = document.createElement('a')
                a.setAttribute('href',"retrieveproswpiko?id="+constdata.anafores[anafora].idperson);
                a.innerHTML = constdata.anafores[anafora].surname+" "+constdata.anafores[anafora].name;
                span1.appendChild(a)
                span2 = document.createElement('span')
                span2.innerHTML = " "+constdata.anafora[anafores].perigrafi
                lielement.appendChild(span1)
                lielement.appendChild(span2)
                ulelement.appendChild(lielement)
        }
    }
    $("#datatoggle").click(function(e){
        if (state == 0){
            state = 1
            parseDataFull(constdata)
        }else{
            state = 0
            parseDataMin(constdata)
        }
    });
};

function parseDataMin(data){
    $("#result").html("")
    $("#result").append("<button id=\"datatoggle\">Αναλυτική</button>")
    if (constdata.proswpiko == null){
        $("#result").append("<h3>Κανείς παρών</h3>")
    }else{
        $("#result").append("<h3>Παρόντες</h3>")
        unique_ranks = getUniqueValuesOfKey(data.proswpiko,"rank")
        for (var index in unique_ranks){
            var temp = constdata.proswpiko.filter(obj => {
                return obj.rank === unique_ranks[index]
              })
            $("#result").append("<h4>"+unique_ranks[index]+" "+temp.length+"</h4>")
            }
    }
    if(data.metaboles == null){
        $("#result").append("<h3>Καμία μεταβολή</h3>")
    }else{
        $("#result").append("<h3>Μεταβολές</h3>")
        unique_categories = getUniqueValuesOfKey(data.metaboles,"Category")
        for (var index in unique_categories){
            var temp = data.metaboles.filter(obj => {
                return obj.Category === unique_categories[index]
              })
              $("#result").append("<h4>"+unique_categories[index]+" "+temp.length+"</h4>")
        }
    }
    if(constdata.ypiresies == null){
        $("#result").append("<h3>Καμία υπηρεσία</h3>")
    }else{
        $("#result").append("<h3>Υπηρεσίες "+constdata.ypiresies.length+"</h3>")
    }
    if(constdata.aitiseis == null){
        $("#result").append("<h3>Καμία αίτηση</h3>")
    }else{
        $("#result").append("<h3>Αιτήσεις "+constdata.aitiseis.length+"</h3>")
    }
    if(constdata.anafores == null){
        $("#result").append("<h3>Καμία αναφορά</h3>")
    }else{
        $("#result").append("<h3>Αναφορές not yet implemented</h3>")
    }
    $("#datatoggle").click(function(e){
        if (state == 0){
            state = 1
            parseDataFull(constdata)
        }else{
            state = 0
            parseDataMin(constdata)
        }
    });
}

function getUniqueValuesOfKey(array, key){
    return array.reduce(function(carry, item){
      if(item[key] && !~carry.indexOf(item[key])) carry.push(item[key]);
      return carry;
    }, []);
  }