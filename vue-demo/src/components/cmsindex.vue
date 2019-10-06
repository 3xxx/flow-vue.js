<template>
  <div><!--  type="card" ref="cmsindex" :height="cmsindexHeight"id="el-carousel"-->
    <el-carousel :interval="4000"><!-- height="1800px" -->
      <el-carousel-item v-for="(item,index) in posts" :key="item.id" style="text-align:center">
        <img :src="'https://zsj.itdos.com'+item.imgUrl" width="100%"><!-- 这里要用height=100% -->
      </el-carousel-item>
    </el-carousel>

  <el-row style="margin-top:2%;margin-bottom:2%">
    <!-- <el-col :span="24" style="text-align: left;"> -->
      <div class="grid-content bg-purple-light"><h1>最新消息</h1></div>
    <!-- </el-col> -->
  </el-row>
  <!-- Page Content-->
  <div class="content">
    <div class="decoration"></div>
    <div class="container no-bottom">
      <!-- Page Columns-->
      <cus v-for="(post,index) in posts" :key="post.id">
        <div v-if="!even1(index)" class="decoration hide-if-responsive"></div>
        <div v-if="even(index)" class="one-half-responsive">
          <p class="thumb-left no-bottom">
            <img v-if="post.imgUrl" :src="'https://zsj.itdos.com'+post.imgUrl" alt="img">
            <strong>{{post.title}}</strong>
            <em v-html=getSimpleText(post.html) style="font-size:16px;"><br></em> 
          </p>
          <div class="thumb-clear"></div>
        </div>

        <div v-else-if="!even(index)" class="one-half-responsive last-column">
          <p class="thumb-right no-bottom">
            <img v-if="post.imgUrl" :src="'https://zsj.itdos.com'+post.imgUrl" alt="img">
            <strong>{{post.title}}</strong>
            <em v-html=getSimpleText(post.html) style="font-size:16px;"><br></em>
          </p>
          <div class="thumb-clear" style="margin-top:1%;margin-bottom:1%"></div>
        </div> 
      </cus>
    </div>

    <div class="decoration" style="margin-top:1%;margin-bottom:1%"></div>
  </div>

  <el-row style="margin-top:1%;margin-bottom:1%">
    <el-col :span="24" style="text-align: left;">
      <div class="grid-content bg-purple-light"><h1>人员简介</h1></div>
    </el-col>
  </el-row>

<div class="swiper-container1" id="swiper1">
  <div class="swiper-wrapper1">
      <swiper1 :options="swiperOption">
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/1s.jpg" alt="img">
                    <h4>John Doe</h4>
                    <em>总工程师</em>
                    <strong style="font-size:16px;">我们喜欢研究水, 并且致力于解决区域性缺水问题，从节约能源、生态、环保、绿色、智慧角度，平衡区域水治理!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple-light">
              <div class="content">
                <div class="container"> 
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/2s.jpg" alt="img">
                    <h4>Johanna Pear</h4>
                    <em>业务经理</em>
                    <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Mail</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/3s.jpg" alt="img">
                    <h4>Peter Wave</h4>
                    <em>项目经理</em>
                    <strong style="font-size:16px;">在全球范围内寻求合作，将最好的理念和技术，运用到供水工程中，帮助客户处理系统性生态环境事务，引导社会大众积极参与!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/4s.jpg" alt="img">
                    <h4>Susson vely</h4>
                    <em>规划师</em>
                    <strong style="font-size:16px;">无论从时空还是经济角度，水资源都是十分宝贵的，周而复始，流水不腐户枢不蠹，将生态水充分利用，改善人居环境是我们的追求!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>Bili One</h4>
                    <em>地质工程师</em>
                    <strong style="font-size:16px;">珠三角区域地质整体稳定，为深层隧洞施工提供了良好地质环境，但局部涌水、断层、软弱夹层等在盾构施工中需要引起重视!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/6s.jpg" alt="img">
                    <h4>Susson vely</h4>
                    <em>咨询师</em>
                    <strong style="font-size:16px;">从第三方角度把控整个项目的全局，提供深一层思维，利用先进的BIM和GIS技术和工具，可视化施工和监控整个项目的建设!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>Slide 7 虚位以待</swiper-slide>
        <swiper-slide>Slide 8 虚位以待</swiper-slide>
        <swiper-slide>Slide 9 虚位以待</swiper-slide>
        <swiper-slide>Slide 10 虚位以待</swiper-slide>
        <div class="swiper-pagination" slot="pagination"></div>
        <div class="swiper-button-prev" slot="button-prev"></div>
        <div class="swiper-button-next" slot="button-next"></div>
      </swiper1>
    </div>
  </div>

      <!-- <swiper :options="swiperOption" v-for="i in 2" :key="i" style="margin-bottom: 1rem;">
        <swiper-slide v-for="i in 9" :key="i">Slide {{ i + 1 }}</swiper-slide>
        <div class="swiper-pagination" slot="pagination"></div>
      </swiper> -->
  <el-row style="margin-top:1%;margin-bottom:1%">
    <el-col :span="24" style="text-align: left;">
      <div class="grid-content bg-purple-light"><h1>施工标段</h1></div>
    </el-col>
  </el-row>
<div class="swiper-container2">
  <div class="swiper-wrapper2">
      <swiper2 :options="swiperOption2">
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/8s.jpg" alt="img">
                    <h4>A1标</h4>
                    <em>鲤鱼洲泵站</em>
                    <strong style="font-size:16px;">包括取水口、泵站、输水管和高位水池……</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A1标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/1s.jpg" alt="img">
                    <h4>B1标</h4>
                    <em>高新沙水库和高新沙泵站</em>
                    <strong style="font-size:16px;">水库总库容482万方，大坝长3.77km，高新沙泵站装机72000kw</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">B1标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/2s.jpg" alt="img">
                    <h4>A2标</h4>
                    <em>跨西江和大金山</em>
                    <strong style="font-size:16px;">双线盾构隧洞，2个区间，盾构直径6.28m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A2标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/3s.jpg" alt="img">
                    <h4>B2标</h4>
                    <em>鱼窝头大道和南沙支线</em>
                    <strong style="font-size:16px;">4个预应力内衬区间，干线盾构直径8.3m，南沙支线盾构直径4.1m，内衬直径2.9m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">B2标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/4s.jpg" alt="img">
                    <h4>A3标</h4>
                    <em>龙洲公路</em>
                    <strong style="font-size:16px;">4个钢管内衬区间，盾构直径6.28m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A3标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>B3标</h4>
                    <em>跨广深港高铁和狮子洋</em>
                    <strong style="font-size:16px;">4个预应力内衬区间，盾构直径8.3m，预应力内衬直径6.4m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">B3标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>A4标</h4>
                    <em>广州南二环绕城高速段</em>
                    <strong style="font-size:16px;">4个钢管内衬区间，盾构直径6.28m，钢管直径4.8m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A4标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>B4标</h4>
                    <em>东莞段</em>
                    <strong style="font-size:16px;">4个预应力内衬区间，盾构直径8.3m，预应力内衬直径6.4m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">B4标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>A5标</h4>
                    <em>伦教横八路段</em>
                    <strong style="font-size:16px;">4个钢管内衬区间，盾构直径6.28m，钢管直径4.8m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A5标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>C1标</h4>
                    <em>穿东莞大岭山TBM隧洞</em>
                    <strong style="font-size:16px;">钻爆隧洞、TBM隧洞和颜屋倒虹吸盾构隧洞</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">C1标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>A6标</h4>
                    <em>南沙段盾构隧洞</em>
                    <strong style="font-size:16px;">4个钢管内衬区间，盾构直径6.28m，钢管直径4.8m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A6标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>C2标</h4>
                    <em>杨屋段钻爆隧洞</em>
                    <strong style="font-size:16px;">3.16km长钻爆隧洞</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">C2标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>A7标</h4>
                    <em>跨西樵水道</em>
                    <strong style="font-size:16px;">双线4个钢管内衬区间，盾构直径6.28m，钢管内衬直径4.8m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">A7标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>D1标</h4>
                    <em>罗田泵站</em>
                    <strong style="font-size:16px;">取水口、前池和进水池、压力箱和输水钢管……</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">D1标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>XX标</h4>
                    <em>东莞分干线</em>
                    <strong style="font-size:16px;">3.6km长隧洞和顶管，钻爆隧洞长1.407km，顶管长2.144km，直径3.0m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">D2标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <swiper-slide>
          <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/5s.jpg" alt="img">
                    <h4>D2标</h4>
                    <em>深圳分干线钻爆隧洞和盾构隧洞</em>
                    <strong style="font-size:16px;">11.9km长隧洞，钻爆隧洞长7km，盾构隧洞长4.9km，直径6.28m，内衬直径4.8m</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">D2标</el-button>
                  </div>
                </div>
              </div>
            </div>
        </swiper-slide>
        <div class="swiper-pagination" slot="pagination"></div>
        <div class="swiper-button-prev" slot="button-prev"></div>
        <div class="swiper-button-next" slot="button-next"></div>
      </swiper2>
    </div>
  </div>
  <!-- <el-carousel :interval="4000" height="500px">
    <el-carousel-item v-for="item in 6" :key="item">
      <h3>{{ item }}</h3>
    </el-carousel-item>
  </el-carousel> --> 
    <!-- <el-carousel :interval="5000" arrow="always" height="500px">
      <el-carousel-item v-for="(post,index) in posts" :key="post.id">
        <el-row>
          <el-col :span="12">
            <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/1s.jpg" alt="img">
                    <h4>John Doe-{{post.id}}</h4>
                    <em>总工程师</em>
                    <strong style="font-size:16px;">我们喜欢研究水, 并且致力于解决区域性缺水问题，从节约能源、生态、环保、绿色、智慧角度，平衡区域水治理!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content bg-purple-light">
              <div class="content">
                <div class="container"> 
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/2s.jpg" alt="img">
                    <h4>Johanna Pear</h4>
                    <em>业务经理</em>
                    <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Mail</el-button>
                  </div>
                </div>
              </div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/3s.jpg" alt="img">
                    <h4>Peter Wave</h4>
                    <em>项目经理</em>
                    <strong style="font-size:16px;">在全球范围内寻求合作，将最好的理念和技术，运用到供水工程中，帮助客户处理系统性生态环境事务，引导社会大众积极参与!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="https://zsj.itdos.com/static/images/4s.jpg" alt="img">
                    <h4>Susson vely</h4>
                    <em>规划师</em>
                    <strong style="font-size:16px;">无论从时空还是经济角度，水资源都是十分宝贵的，周而复始，流水不腐户枢不蠹，将生态水充分利用，改善人居环境是我们的追求!</strong>
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-carousel-item>
    </el-carousel> -->
  </div>
</template>

<script type="text/javascript">
  /* eslint-disable */
  const axios = require('axios');
  import { swiper, swiperSlide } from 'vue-awesome-swiper';

  export default { // 这里需要将模块引出，可在其他地方使用
    name: 'cmsindex',
    // components:{
    //   cus
    // },
    components:{
      swiper1:swiper,
      swiper2:swiper
    },
    data:{
        users:{
            name:'',
            age:''
        },
        message: 'Hello Vue!',
        info: null,
        superurl: [
          {
              id:'1',
              url: '',
              img: 'https://zsj.itdos.com/static/img/100126522.jpg',
          },
          {
              id:'2',
              url: '',
              img: 'https://zsj.itdos.com/static/img/100165705.jpg',
          },
          {
              id:'3',
              url: '',
              img: 'https://zsj.itdos.com/static/img/100268041.jpg',
          },
          {
              id:'4',
              url: '',
              img: 'https://zsj.itdos.com/static/img/16pic_1227574_b.jpg',
          },
          {
              id:'5',
              url: '',
              img: 'https://zsj.itdos.com/static/img/254247.jpg',
          },
          {
              id:'6',
              url: '',
              img: 'https://zsj.itdos.com/static/img/500171023.jpg',
          },
        ],
        posts:[],
      },
      data() {
        var checkName = (rule, value, callback) => {
          if (value === '') {
            callback(new Error('请输入账号'));
          } else {
            callback();
          }
        };
        var checkNum = (rule, value, callback) => {
          if (!value) {
            return callback(new Error('账号不能为空'));
          }
          setTimeout(() => {
            if (!Number.isInteger(value)) {
              callback(new Error('请输入数字值'));
            } else {
              var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
              if (!myreg.test(value) ) {
                callback(new Error('请输入正确的手机号码'));
              } else {
                callback();
              }
            }
          }, 1000);
        };
        var validatePass = (rule, value, callback) => {
          if (value === '') {
            callback(new Error('请输入密码'));
          } else {
            callback();
          }
        };
        return {
          total: 50,
          currentPage: 1,
          pageSize: 10,
          loginPower:false,
          checked: true,
          /*插入form方法*/
          /*设定规则指向*/
          ruleForm2: {
            name:'',
            pass: '',
            num: '',
            delivery: false,
          },
          rules2: {
            typename: [
              { validator: checkName, trigger: 'blur' }
            ],
            pass: [
              { validator: validatePass, trigger: 'blur' }
            ],
            num: [
              { validator: checkNum, trigger: 'blur' }
            ]
          },
          /*插入form方法*/
          dialogTableVisible: false,
          dialogFormVisible: false,
          form: {
            name: '',
            type: [],
            resource: '',
            desc: ''
          },
          formLabelWidth: '120px',

          isCollapse: true,
          bannerHeight:200,

          cmsindexHeight:200,
          // calleft:0
          posts:[],
          numbers:0,

          // dialogFormVisible: false,
          form: {
            name: '',
            region: '',
            date1: '',
            date2: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
          },
          search: '',
          superurl: [
              {
                  id:'1',
                  url: '',
                  img: 'https://zsj.itdos.com/static/img/100126522.jpg',
              },
              {
                  id:'2',
                  url: '',
                  img: 'https://zsj.itdos.com/static/img/100165705.jpg',
              },
              {
                  id:'3',
                  url: '',
                  img: 'https://zsj.itdos.com/static/img/100268041.jpg',
              },
              {
                  id:'4',
                  url: '',
                  img: 'https://zsj.itdos.com/static/img/16pic_1227574_b.jpg',
              },
              {
                  id:'5',
                  url: '',
                  img: 'https://zsj.itdos.com/static/img/254247.jpg',
              },
              {
                  id:'6',
                  url: '',
                  img: 'https://zsj.itdos.com/static/img/500171023.jpg',
              },
          ],
          swiperOption: {
            slidesPerView: 3,
            spaceBetween: 30,
            loop: true,
            loopFillGroupWithBlank: true,
            autoplay: {
              delay: 2500,
              disableOnInteraction: false
            },
            pagination: {
              el: '.swiper-pagination',
              clickable: true
            },
            navigation: {
              nextEl: '.swiper-button-next',
              prevEl: '.swiper-button-prev'
            }
          },
          swiperOption2: {
            slidesPerView: 3,
            slidesPerColumn: 2,
            spaceBetween: 30,
            // loop: true,
            // loopFillGroupWithBlank: true,
            pagination: {
              el: '.swiper-pagination',
              clickable: true
            },
            navigation: {
              nextEl: '.swiper-button-next',
              prevEl: '.swiper-button-prev'
            }
          }
          // bannerHeight: 700,
          // screenWidth: 1920,
        };
      },
    mounted:function () {
      this.sendGetByObj();
      // 获取浏览器可视区域高度
      this.cmsindexHeight =   `${document.documentElement.clientHeight}`;
      console.log(this.cmsindexHeight);
      //document.body.clientWidth;
      //console.log(self.clientHeight);
      window.onresize = function temp() {
        this.cmsindexHeight = `${document.documentElement.clientHeight}`;
      };
      // this.setSize1();
      // const that = this;
      // //监听浏览器窗口大小改变
      // window.addEventListener('resize', function() {
      //   var width = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
      //   that.screenWidth = width;
      //   that.setSize();
      // }, false);
    },
    watch: {
      // 如果 `clientHeight` 发生改变，这个函数就会运行
      cmsindexHeight: function () {
        this.changeFixed(this.cmsindexHeight)
      }
    },
      methods:{
        resetForm(formName) {
          this.$refs[formName].resetFields();
        },

        changeWidthL(key, keyPath) {
          console.log(key, keyPath);
        },

        handleOpen(key, keyPath) {
          console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
          console.log(key, keyPath);
        },
        handleEdit(index, row) {
          console.log(index, row);
        },
        handleDelete(index, row) {
          console.log(index, row);
        },
        deleteRow(index, rows) {//删除改行
          rows.splice(index, 1);
        },
        addRow(docstatedata,event){
          docstatedata.push({ ID:'', Name: ''})
        },

        //html剔除富文本标签，留下纯文本
        getSimpleText(html){
          var re1 = new RegExp("<.+?>","g");//匹配html标签的正则表达式，"g"是搜索匹配多个符合的内容
          var msg = html.replace(re1,'');//执行替换成空字符
          return '<br>'+msg.substring(0,100)+'……';
        },
        even: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers % 2 === 0
          // })
        },
        even1: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers === 0
          // })
        },
        changeCollapse:function(isCollapse){
          if (isCollapse==true) {
            this.isCollapse=false
          }else{
            this.isCollapse=true
          }
        },
        sendGetByObj(currentPage){
          axios({
            method: 'get',
            url: 'https://zsj.itdos.com/v1/wx/getlistarticles',//2.get通过params选项
            params:{
              page:1
            }
          })
          .then(response => (this.posts = response.data.articles))
          .catch(function (error) {
            console.log(error);
          });
        },
        toNext: function(index) {
          sessionStorage.ticketName =this.sortList[index].list;
          this.$router.push('/mine/tiketOrder');
        },
        handleSizeChange: function (size) {
          this.pageSize = size;
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.docstate(currentPage);
        },
        // changeFixed(cmsindexHeight){ //动态修改样式
        //   // console.log(cmsindexHeight);
        //   // console.log(this.$refs.cmsindex);
        //   // console.log(this.$refs.cmsindex.$el.style.height);
        //   this.$refs.cmsindex.$el.style.height = cmsindexHeight-20+'px';
        // },
      // setSize1: function() {
      //   var width = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
      //   this.screenWidth = width;
      //   //图片                高 / 宽  700 / 1920
      //   this.bannerHeight = 700 / 1920 * this.screenWidth - 50
      //   document.getElementById('el-carousel').style.height = this.bannerHeight + 'px';
 
      // },
      // setSize: function() {
      //   this.bannerHeight = 700 / 1920 * this.screenWidth - 50
      //   document.getElementById('el-carousel').style.height = this.bannerHeight + 'px';
      // },
    }
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/*  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
  .el-carousel-item {
    text-align:center 放这里无效，为何
  }*/

/*  .el-carousel__container {
    height: 100% !important;
  }
  .el-carousel img {
    display: inline-block;
    height: auto;
    max-width: 100%;
  }*/
  /*swiper轮播高度*/
  .swiper-container1 {
    width: 100%;
    height: 500px;
  }

  .swiper-container2 {
    height: auto!important;
    margin-left: auto;
    margin-right: auto;
  }
  .swiper-container2 .swiper-slide {
    height: 400px;
  }

  #vue{
    color: green;
    font-size: 28px;
  }
  .con_section{
    position: absolute;
    top: 0px;
    bottom: 0px;
    left:0px;
    width:100%;
  }
  .blueheader {
    height: 60px;
    line-height: 60px;
    background: #67c23a;
    color: #fff;
  }
  .el-menu-item.is-active {
    color: #67c23a;
  }
  .headlogo{
    float: left;
    height: 60px;
    margin: 0 20px;
    width: 300px;
  }
  ul.el-menu {
    background: #e4e8f1;
  }
  .userinfo{
    position: absolute;
    right: 0;
  }
  .el-submenu__title{
    background:#eef1f6;
  }
  .el_main{
    padding:0px;
  }
/*  .home_main{
    padding:10px;
  }*/
  .breadcrumb-container .title {
    width: 200px;
    float: left;
    color: #475669;
    font-size: 13px;
  }
  .breadcrumb-inner {
      float: right;
      font-size: 13px;
  }
  .el-breadcrumb__inner, .el-breadcrumb__inner a {
    font-weight: 400;
  }

  /*正文*/
  .overlay{
    pointer-events:none;
    position:absolute;
    width:100%;
    height:100%;
    background-color:rgba(0,0,0,0.5);
    z-index:9999;
  }
  
  .content{
    clear:both;
    margin-left:30px;
    margin-right:30px;
  }
  @media (min-width:768px){
    .content{
      margin-left:70px;
      margin-right:70px;  
    }
  }
  
  .decoration{ 
    height:1px;
    background-color:rgba(0,0,0,0.1);
    margin-bottom:30px;
    display:block;
    clear:both;
  }
  
  .container{
    margin-bottom:30px;
  }
  .no-bottom{
    margin-bottom:0px;
    padding-bottom:0px;
  }
  
  @media (min-width:760px) { 
    .one-third-responsive{
      width:30%;
      float:left;
      margin-right:5%;  
    }
    
    .one-half-responsive{
      width:46%;
      float:left;
      margin-right:8%;
    }
    
    .sidebar-left-big{
      width:70%;
      float:left;
      margin-right:5% 
    }
    
    .sidebar-right-small{
      width:25%;
      float:right;  
    }
      
    .sidebar-right-big{
      width:70%;
      float:right;  
    }
    
    .sidebar-left-small{
      width:25%;
      float:left;
      margin-right:5%;  
    }
    
    .hide-if-responsive{
      display:none;
    }
  }
  
  /*Thumbnails Columns*/
  .thumb-clear{
    height:40px;
    display:block;
    width:100%;
  }
  .thumb-left{
    line-height:24px;
    display:block;
    padding-bottom:10px;
  }
  .thumb-left a{
    display:block;
    text-align:right;
  }
  .thumb-left img{
    width:100px;
    height:100px;
    border-radius:100px;
    float:left;
    margin-right:20px;
      transition:all 500ms ease;
  }
  .thumb-left img:hover{
      transform:scale(0.9, 0.9);
      transition:all 500ms ease;
  }
  .thumb-left strong{
    color:#1a1a1a;
    display:inline-block;
    padding-bottom:5px;
    font-size:13px;
  }
  .thumb-left em{
    font-style:normal;
  }
  .thumb-right{
    line-height:24px;
    display:block;
    padding-bottom:10px;
  }
  .thumb-right img{
    width:100px;
    height:100px;
    border-radius:100px;
    float:right;
    margin-left:20px;
      transition:all 500ms ease;
  }
  .thumb-right img:hover{
      transform:scale(0.9, 0.9);
      transition:all 500ms ease;
  }
  .thumb-right strong{
    color:#1a1a1a;
    display:inline-block;
    padding-top:5px;
    padding-bottom:5px;
    font-size:13px;
  }
  .thumb-right em{
    font-style:normal;
  }
  @media (min-width:600px){
    .thumb-left img{
      width:140px;
      height:140px;
      border-radius:150px;  
    }
    
    .thumb-left em{
      line-height:28px; 
    }
    
    .thumb-left strong{
      padding-top:10px; 
    }
    
    .thumb-right img{
      width:140px;
      height:140px;
      border-radius:150px;  
    }
    
    .thumb-right em{
      line-height:28px; 
    }
    
    .thumb-right strong{
      padding-top:10px; 
    } 
  }
  
  .last-column{
    margin-right:0%!important;
  }

  .prev-staff, .prev-quote{
    z-index:99999;
    /*background-image:url(../ky_img/prev1.png);*/
    background-repeat:no-repeat;
    width:45px;
    height:45px;
    background-size:16px 16px;
      background-position: 13px 15px;
    position:absolute;
    left:0px;
    margin-top:20%; 
    border-radius:45px;
    transition:all 400ms ease;
  }
  /*
  *   Owl Carousel Owl  Theme 
  * v1.24
  */
  .slider-controls{
    margin-bottom:-23px;
  }
  .third-thumb{
    width:29%;
    float:left;
    margin-left:2%;
    margin-right:2%;
  }
  .half-thumb{
    width:48%;
    float:left;
    margin-left:1%;
    margin-right:1%;
  }
  .next-slider{
    /*background-image:url(../ky_img/next.png);*/
    background-repeat:no-repeat;
    width:27px;
    height:27px;
    background-size:27px;
    height:27px;
    position:absolute;
    right:20px;
    margin-top:-45%;
  }
  .prev-slider{
    /*background-image:url(../ky_img/prev.png);*/
    background-repeat:no-repeat;
    width:27px;
    height:27px;
    background-size:27px;
    height:27px;
    position:absolute;
    left:20px;
    margin-top:-45%;
  }
  .prev-staff, .prev-quote{
    z-index:99999;
    /*background-image:url(/static/index/ky_img/prev1.png);*/
    background-repeat:no-repeat;
    width:45px;
    height:45px;
    background-size:16px 16px;
      background-position: 13px 15px;
    position:absolute;
    left:0px;
    margin-top:20%; 
      border-radius:45px;
      transition:all 400ms ease;
  }
  .prev-staff:hover, .prev-quote:hover{
      background-color:rgba(0,0,0,0.1);
      transition:all 400ms ease;
  }
  .next-staff, .next-quote{
    z-index:99999;
    /*background-image:url(/static/index/ky_img/next1.png);*/
    background-repeat:no-repeat;
    width:45px;
    height:45px;
    background-size:16px 16px;
      background-position: 16px 15px;
    position:absolute;
    right:0px;
    margin-top:20%;   
      border-radius:45px;
      transition:all 400ms ease;
  }
  .next-staff:hover, .next-quote:hover{
       background-color:rgba(0,0,0,0.1);
      transition:all 400ms ease;   
  }
  @media screen and (orientation:landscape) {
    .next-staff{
      margin-top:24%;
    }
    
    .prev-staff{
      margin-top:24%; 
    }
    
    .next-quote{
      margin-top:4%;  
    }
    
    .prev-quote{
      margin-top:4%;  
    }
    
    .staff-item strong{
      width:180px;
      margin-left:auto;
      margin-right:auto;
    }
  }
  @media only screen 
  and (min-device-width : 768px) 
  and (max-device-width : 1024px) 
  and (orientation : landscape) { 
    .next-staff{
      margin-top:12.1%;
    }
    .prev-staff{
      margin-top:12.1%; 
    }
    
    .next-quote{
      margin-top:4%;  
    }
    
    .prev-quote{
      margin-top:4%;  
    }
    
    .staff-item strong{
      width:180px;
      margin-left:auto;
      margin-right:auto;
    }
  }
  @media only screen 
  and (min-device-width : 768px) 
  and (max-device-width : 1024px) 
  and (orientation : portrait) { 
    .next-staff{
      margin-top:17.3%;
    }
    .prev-staff{
      margin-top:17.3%; 
    }
    
    .next-quote{
      margin-top:8%;  
    }
    
    .prev-quote{
      margin-top:8%;  
    }
    
    .staff-item strong{
      width:180px;
      margin-left:auto;
      margin-right:auto;
    }
  }
  
  /*////////////////////*/
  /*Quote & Staff Slider*/
  /*////////////////////*/
  .staff-item img{
    border-radius:300px;  
    max-width:150px;
    margin-left:auto;
    margin-right:auto;
    transition:all 500ms ease;
    margin-top:10px;
    text-align:center;
    /*width: 400px;*/
    /*height: 300px;*/
    /*border: 1px dashed #ccc;*/
    display: table-cell; /*//主要是这个属性*/
    /*vertical-align: middle;*/
    /*text-align: center;*/
  }
  .staff-item img:hover{
      transform:scale(1.1, 1.1);
      transition:all 500ms ease;
  }
  .staff-item h4{
    text-align:center;
    margin-top:20px;  
    font-size:16px;
  }
  .staff-item em{
    text-align:center;
    font-size:12px;
    color:#e34e47;
    display:block;
    margin-top:-10px;
    margin-bottom:10px;
    font-style:normal;
  }
  .staff-item strong{
    display:block;  
    font-weight:400;
    text-align:center;
    padding-left:20px;
    padding-right:20px;
    line-height:28px;
  }
  /* 8. Quote Slider */
  .quote-slider h4{
    font-family:'Source Sans Pro', sans-serif;
    font-size:18px;
    font-weight:300;  
    width:87%;
    text-align:center;
    margin-left:auto;
    margin-right:auto;
    line-height:36px;
  }
  .quote-slider a{
    text-align:center;
    margin-top:20px;  
  }
</style>
