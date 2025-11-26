document.querySelectorAll('.slide').forEach(function(slide) {
  slide.style.position = 'absolute';
  slide.style.top = 0;
  slide.style.left = '100%';
});

document.querySelectorAll('.slide.hidden').forEach(function(slide) {
  slide.classList.remove('hidden');
  slide.style.opacity = 0;
});

slideAnimationTime = 2500;
easingType = 'easeInOutCubic';
easingType = 'linear';

updateSlideVisibility = function () {
  if (previousPage < 0) {
    getSlideElements().forEach(function(slide) {
      if (parseInt(slide.id.slice(6)) == currentPage) {
        slide.style.opacity = 1;
        slide.style.left = '0';
        slide.classList.add('visible')
      }
    });
    return;
  }
  getSlideElements().forEach(function(slide) {
    if (parseInt(slide.id.slice(6)) != currentPage && slide.classList.contains('visible')) {
      //slide.classList.add('hidden');
      slide.classList.remove('visible');
      animate(document.querySelector('#slide-' + parseInt(slide.id.slice(6))), {
        opacity: 0,
        duration: slideAnimationTime,
        ease: easingType, // 'easing' is now 'ease'
        onFinish: function(anim) {
          slide.style.left = '100%';
        }
      });


    }
  });
  getSlideElements().forEach(function(slide) {
    if (parseInt(slide.id.slice(6)) == currentPage) {
      slide.style.left = '0';
      slide.classList.remove('hidden');
      slide.classList.add('visible');
      animate(document.querySelector('#slide-' + currentPage), {
        opacity: 1,
        duration: slideAnimationTime,
        ease: easingType // 'easing' is now 'ease'
      });
    }
  });
}
