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
    $("#result").append("<button id=\"datatoggle\">Μερική</button>")
    $("#result").append("<h3>Παρόντες</h3>")
    for (var rank in data.rankmap.map) {
        if (data.rankmap.map[rank]){
            $("#result").append("<h3>"+rank+" "+data.rankmap.map[rank].length+"</h3>")
            $("#result").append("<ul>")
            if (data.rankmap.map.hasOwnProperty(rank)) {           
                for (var person in data.rankmap.map[rank]){
                    $("#result").append("<li>"+data.rankmap.map[rank][person].surname+" "+data.rankmap.map[rank][person].name+"</li>")        
                }
            }
            $("#result").append("</ul>")
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
    $("#result").append("<h3>Παρόντες</h3>")
    for (var rank in data.rankmap.map) {
        if (data.rankmap.map[rank]){
            $("#result").append("<h4>"+rank+" "+data.rankmap.map[rank].length+"</h4>")
        }
    }
    if(data.metaboles == null){
        $("#result").append("<h3>Καμία μεταβολή</h3>")
    }else{
        $("#result").append("<h3>Μεταβολές</h3>")
        unique_categories = getUniqueValuesOfKey(data.metabolesmin,"category")
        for (var index in unique_categories){
            $("#result").append("<h4>"+unique_categories[index]+"</h4>")
            ulelement = document.createElement('ul')
            var temp = data.metabolesmin.filter(obj => {
                return obj.category === unique_categories[index]
              })
              for (var i in temp){
                  lielement = document.createElement('li')
                  lielement.innerHTML = temp[i].rank +" "+temp[i].count
                  ulelement.appendChild(lielement)
              }
              $("#result").append(ulelement)
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
}

function getUniqueValuesOfKey(array, key){
    return array.reduce(function(carry, item){
      if(item[key] && !~carry.indexOf(item[key])) carry.push(item[key]);
      return carry;
    }, []);
  }