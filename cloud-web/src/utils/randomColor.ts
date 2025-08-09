const generateRandomHexColor = ( ): string => {
  const maxVal = 0xFFFFFF;
  let randomNumber = Math.floor(Math.random() * (maxVal + 1));
  let hexColor = randomNumber.toString(16);
  hexColor = hexColor.padStart(6, '0');
  return `#${hexColor.toUpperCase()}`;
}



export default generateRandomHexColor;
