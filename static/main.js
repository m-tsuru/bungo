const main = document.querySelector('#main');

if(main) {
    // gsap.registerPlugin(ScrollTrigger); // npm/yarnの際に必要
    const panels = gsap.utils.toArray('.panel');
    const wrapperWidth = wrapper.offsetWidth;
    /**
    * 横スクロール開始
    */
    gsap.to( panels, {
        xPercent: -100 * (panels.length - 1), // transformX
        ease: "none", // easingの設定
        scrollTrigger: { // scrollTrigger
            trigger: wrapper, // アニメーションの対象となる要素
            pin: true, // 要素を固定する
            scrub: 1, // スクロールとアニメーションを同期させる。数値で秒数の設定に
            snap: { // スナップスクロールにする
                snapTo: 1 / ( panels.length - 1 ), // スナップで移動させる位置
                duration: {min: .4, max: .6}, // スナップで移動する際の遅延時間
                ease: "none" // easing
            },
            end: () => "+=" + wrapperWidth // アニメーションの終了タイミング
        }
    })
}
