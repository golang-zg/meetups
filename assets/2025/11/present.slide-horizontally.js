document.querySelectorAll('.slide').forEach(function(slide) {
  slide.style.position = 'absolute';
  slide.style.top = 0;
  slide.style.left = '100%';
});

document.querySelectorAll('.slide.hidden').forEach(function(slide) {
  slide.classList.remove('hidden');
  slide.style.opacity = 0;
});

slideAnimationTime = 1500;
easingType = 'easeInOutCubic';
easingType = 'linear';
firstAnimation = true

updateSlideVisibility = function () {
  if (firstAnimation) {
    firstAnimation = false;
    getSlideElements().forEach(function(slide) {
      if (parseInt(slide.id.slice(6)) < currentPage) {
        slide.style.opacity = 1;
        slide.style.left = -1 * window.innerWidth + 'px';
        slide.classList.add('visible')
        //slide.style.transform = 'translateX(-'+window.innerWidth+'px)';
      }
      if (parseInt(slide.id.slice(6)) == currentPage) {
        slide.style.opacity = 1;
        slide.style.left = 0;
        slide.classList.add('visible')
        animate(slide, {
          //left: '0',
          opacity: anime.linear(0,1),
          scale: 1,
          left : 0,
          duration: slideAnimationTime,
        });
        //slide.style.transform = 'translateX(0px)';
      } else {
        direction = 1
        if (parseInt(slide.id.slice(6)) > currentPage) {
          direction = -1
        }
        animate(slide, {
          //left: '0',
          opacity: anime.linear(1,"0.9 95%", 0),
          scale: anime.linear(1,0),
          left : anime.linear(0, window.innerWidth * direction * -1),
          duration: slideAnimationTime,
        });
      }
    });
    return;
  }
  direction = 1
  if (currentPage < previousPage) {
    direction = -1
  }
  if (currentPage == previousPage) {
    console.log('same page')
     return;
  }
  if (previousPage < 0 && !firstAnimation) {
    previousPage = currentPage + direction
  }
  console.log('Animating from', previousPage, 'to', currentPage, 'direction', direction)
  getSlideElements().forEach(function(slide) {
    let id = slide
    if (parseInt(slide.id.slice(6)) == previousPage) {
      animate(id, {
        //left: '0',
        opacity: anime.linear(1,"0.9 95%", 0),
        scale: {
          from: 1,
          to: 0.5,
          easing: 'easeInOutCubic'
        },
        left : {
          from: 0,
          to: window.innerWidth * direction * -1
        },
        duration: slideAnimationTime,
        delay: slideAnimationTime / 15
      });
    }
    if (parseInt(slide.id.slice(6)) == currentPage){
      animate(id, {
        left: anime.linear(window.innerWidth * direction, 0),
        opacity: {
          from: 0,
          to: 1,
        },
        scale: {
          from: 1,
          to: 1
        },
        duration: slideAnimationTime
      })
    }
  });
}
