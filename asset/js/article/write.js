var vue = new Vue(
    {
        el:"#page",
        data:{
            articles:[],
            channels:[]
        },
        methods:{
                editarticle:function(item){

                },
                loadchannel:function(){
                        var self = this;
                        restgo.post("channel/search",{deleted:0}).then(function(res){
                            self.channels = res.rows;
                        })
                },
                loadarticle:function(chid){

                }
        },
        created:function(){
                this.loadchannel();
                this.loadarticle();
        },
        mounted:function(){

        }
    }
)