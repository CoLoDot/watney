type NeoObject = {
  closeApproachDate: string;
  estimatedDiameterInMeters: number;
  id: string;
  name: string;
  nasaJplUrl: string;
  relativeVelocityKMPerSecond: number;
};

export type NeoResponse = {
  totalObjects: number;
  date: string;
  nearEarthObjects: NeoObject[];
};
